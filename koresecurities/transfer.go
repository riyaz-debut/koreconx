package koresecurities

import (
	"encoding/json"
	"fmt"

	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/person"
	"kore_chaincode/user"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// TransferSecurities saves the SecuritiesTransferRequest hash in world state
func TransferSecurities(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(SecuritiesTransferRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeTransfer
	request.Status = 0
	request.UpdatedAt = request.CreatedAt

	// Company
	company, err := user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// Requestor associated?
	found, err := company.IsRequestorAssociated(request.TransferRequestor)
	if !found {
		return nil, err
	}

	// TA Associated?
	if request.TransferApprover != "" {
		found = company.IsRequestorTA(request.TransferApprover)
		if !found {
			return nil, status.ErrInternal.WithMessage("Transfer Approver is not associated with the company")
		}
	}

	// securitiy id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// owner id
	_, err = person.GetPersonByID(ctx, request.OwnerID)
	if err != nil {
		return nil, err
	}

	// transferred to id
	transferTo, err := person.GetPersonByID(ctx, request.TransferredToID)
	if err != nil {
		return nil, err
	}

	_, err = transferTo.IsKYCVerification(request.EffectiveDate)
	if err != nil {
		return nil, err
	}

	holding, _, err := GetHoldingBySecuritiesID(ctx, request.OwnerID, request.KoresecuritiesID)
	if err != nil {
		return nil, err
	}

	if holding.AvailableShares < request.TotalSecurities {
		return nil, status.ErrInternal.WithMessage("Owner does not have enough securities to transfer.")
	}

	// change strcut to json
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(ctx.GetStub().GetTxID(), jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseID)
	response.ID = ctx.GetStub().GetTxID()
	return response, nil
}

// UpdateTransferSecurities saves the SecuritiesTransferRequest hash in world state
func UpdateTransferSecurities(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
	request := new(UpdateSecuritiesTransferRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// Company
	transferRequest, err := GetSecuritiesTransferRequestByID(ctx, request.TransferRequestID)
	if err != nil {
		return nil, err
	}

	if transferRequest.Status != 0 {
		return nil, status.ErrBadRequest.WithMessage("Transfer request must be a pending request!")
	}

	// Company
	_, err = user.GetCompanyByID(ctx, transferRequest.CompanyID)
	if err != nil {
		return nil, err
	}

	transferRequest.UpdatedAt = request.CreatedAt

	// Reject the request
	if request.Status == 2 {
		transferRequest.Status = 2
		transferRequest.Reason = request.Reason
		jsonData, err := json.Marshal(transferRequest)
		if err != nil {
			return nil, status.ErrInternal.WithError(err)
		}

		err = ctx.GetStub().PutState(request.TransferRequestID, jsonData)
		if err != nil {
			return nil, status.ErrInternal.WithError(err)
		}
		response := new(utils.ResponseMessage)
		response.Message = "The transfer request status has been rejected!"
		return response, nil
	}

	// accept the request
	transferRequest.Status = 1
	holding, holdingID, err := GetHoldingBySecuritiesID(ctx, transferRequest.OwnerID, transferRequest.KoresecuritiesID)
	if err != nil {
		return nil, err
	}

	if holding.AvailableShares < transferRequest.TotalSecurities {
		return nil, status.ErrInternal.WithMessage("Owner does not have enough securities to transfer.")
	}

	holding.AvailableShares = holding.AvailableShares - transferRequest.TotalSecurities
	holding.HoldingAmount = holding.HoldingAmount - transferRequest.TotalSecurities
	holding.UpdatedAt = request.CreatedAt

	// change strcut to json
	jsonData, err := json.Marshal(holding)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(holdingID, jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// change strcut to json
	jsonData, err = json.Marshal(transferRequest)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.TransferRequestID, jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// Get buyers holding
	holding, holdingID, err = GetHoldingBySecuritiesID(ctx, transferRequest.TransferredToID, transferRequest.KoresecuritiesID)

	if holdingID != "" {
		holding.HoldingAmount = holding.HoldingAmount + transferRequest.TotalSecurities
		holding.AvailableShares = holding.AvailableShares + transferRequest.TotalSecurities
		holding.UpdatedAt = request.CreatedAt
	} else {
		holding = Holding{}
		holding.CompanyID = transferRequest.CompanyID
		holding.SecuritiesHolderID = transferRequest.TransferredToID
		holding.KoresecuritiesID = transferRequest.KoresecuritiesID
		holding.HoldingAmount = transferRequest.TotalSecurities
		holding.AvailableShares = transferRequest.TotalSecurities
		holding.CreatedAt = request.CreatedAt
		holding.DocType = utils.DocTypeHolding
		holding.LastUpdatedAt = request.CreatedAt
		holding.UpdatedAt = request.CreatedAt
		holding.ReasonCode = ""
		holdingID = request.TransactionID
	}

	jsonData, err = json.Marshal(holding)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(holdingID, jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseMessage)
	response.Message = "The transfer request status has been approved!"
	return response, nil
}

// GetSecuritiesTransferRequestByID fetches the SecuritiesTransferRequest with the given ID from world state
func GetSecuritiesTransferRequestByID(ctx contractapi.TransactionContextInterface, ID string) (SecuritiesTransferRequest, error) {
	data := SecuritiesTransferRequest{}
	dataBA, err := SecuritiesTransferRequestExists(ctx, ID)
	if err != nil {
		return data, err
	}

	// ummarshal the byte array to structure
	err = json.Unmarshal(dataBA, &data)
	if err != nil {
		return data, status.ErrInternal.WithError(err)
	}
	return data, nil
}

// SecuritiesTransferRequestExists checks whether the SecuritiesTransferRequest with the given ID exists or not
func SecuritiesTransferRequestExists(ctx contractapi.TransactionContextInterface, ID string) ([]byte, error) {
	var queryString string
	queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeTransfer, ID)
	dataBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("Transfer request with ID: %s does not exists!", ID))
	return dataBA, err
}
