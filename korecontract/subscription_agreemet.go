package korecontract

import (
	"encoding/json"
	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/user"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddSubscriptionAgreement saves the subscription agreement hash in world state
func AddSubscriptionAgreement(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(SubscriptionAgreement)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeSubscriptionAgreement
	request.UpdatedAt = request.CreatedAt

	// company
	_, err = user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// document
	if request.DocumentID != "" {
		_, err = GetDocumentByID(ctx, request.DocumentID, request.CompanyID)
		if err != nil {
			return nil, err
		}
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
