package trade

import (
	"encoding/json"
	"fmt"
	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/koresecurities"
	"kore_chaincode/person"
	"kore_chaincode/user"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddTradeRequest adds a new tradeRequest in world state
func AddTradeRequest(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	tradeRequest := new(TradeRequest)

	// change json to strcut
	err := json.Unmarshal(data, tradeRequest)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	tradeRequest.DocType = utils.DocTypeTradeRequest
	tradeRequest.UpdatedAt = tradeRequest.CreatedAt

	// Company
	_, err = user.CompanyExists(ctx, tradeRequest.CompanyID)
	if err != nil {
		return nil, err
	}

	// securitiy id
	_, err = koresecurities.GetSecuritiesByID(ctx, tradeRequest.KoresecuritiesID, tradeRequest.CompanyID)
	if err != nil {
		return nil, err
	}

	// Shareholder id
	_, err = person.GetPersonByID(ctx, tradeRequest.ShareholderID)
	if err != nil {
		return nil, err
	}

	// change strcut to json
	jsonData, err := json.Marshal(tradeRequest)
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

// AddATSTrade adds a new tradeRequest in world state
func AddATSTrade(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(AtsTradeRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// Company
	_, err = user.CompanyExists(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// securitiy id
	_, err = koresecurities.GetSecuritiesByID(ctx, request.KoresecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// owner id
	_, err = person.GetPersonByID(ctx, request.OwnerID)
	if err != nil {
		return nil, err
	}

	// transferred to id
	_, err = person.GetPersonByID(ctx, request.TransferredToID)
	if err != nil {
		return nil, err
	}

	// find the hold transaction
	holdShares, _, err := koresecurities.GetHoldSharesByID(ctx, request.ATSTransactionID, request.OwnerID, request.KoresecuritiesID)
	if err != nil {
		return nil, err
	}

	if holdShares.AvailableNumberOfShares != request.TotalSecurities {
		return nil, status.ErrInternal.WithMessage("Number of shares does not match with the hold share transaction.")
	}

	// Create the ATS trade request
	tradeRequest := AtsTrade{}
	tradeRequest.CompanyID = request.CompanyID
	tradeRequest.KoresecuritiesID = request.KoresecuritiesID
	tradeRequest.RequestorID = request.RequestorID
	tradeRequest.ATSTransactionID = request.ATSTransactionID
	tradeRequest.OwnerID = request.OwnerID
	tradeRequest.TransferredToID = request.TransferredToID
	tradeRequest.AuthorizationID = request.AuthorizationID
	tradeRequest.TotalSecurities = request.TotalSecurities
	tradeRequest.TradePrice = request.TradePrice
	tradeRequest.EffectiveDate = request.EffectiveDate
	tradeRequest.CreatedAt = request.CreatedAt
	tradeRequest.DocType = utils.DocTypeATSTrade
	tradeRequest.UpdatedAt = request.CreatedAt

	// update the hold share request
	holdShares.Status = "1"
	jsonData, err := json.Marshal(holdShares)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.ATSTransactionID, jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// save the ATS trade
	jsonData, err = json.Marshal(tradeRequest)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(ctx.GetStub().GetTxID(), jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// update the seller holdings
	sellerHolding, sellerHoldingID, err := koresecurities.GetHoldingBySecuritiesID(ctx, request.OwnerID, request.KoresecuritiesID)

	if sellerHoldingID == "" {
		return nil, status.ErrInternal.WithMessage("Seller does not hold the shares.")
	}

	sellerHolding.HoldingAmount = sellerHolding.HoldingAmount - request.TotalSecurities
	sellerHolding.UpdatedAt = request.CreatedAt

	jsonData, err = json.Marshal(sellerHolding)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(sellerHoldingID, jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// Get buyers holding
	holding, holdingID, err := koresecurities.GetHoldingBySecuritiesID(ctx, request.TransferredToID, request.KoresecuritiesID)

	if holdingID != "" {
		holding.HoldingAmount = holding.HoldingAmount + request.TotalSecurities
		holding.AvailableShares = holding.AvailableShares + request.TotalSecurities
		holding.UpdatedAt = request.CreatedAt
	} else {
		holding = koresecurities.Holding{}
		holding.CompanyID = request.CompanyID
		holding.SecuritiesHolderID = request.TransferredToID
		holding.KoresecuritiesID = request.KoresecuritiesID
		holding.HoldingAmount = request.TotalSecurities
		holding.AvailableShares = request.TotalSecurities
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

	response := new(utils.ResponseID)
	response.ID = ctx.GetStub().GetTxID()
	return response, nil
}

// GetAllTradeRequests returns all tradeRequests found in world state
func GetAllTradeRequests(ctx contractapi.TransactionContextInterface, data []byte) ([]TradeRequestDoc, error) {
	request := new(AllTradeRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// find the company
	company, err := user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsRequestorAssociated(request.RequestorID)
	if !found {
		return nil, err
	}

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"shareholder_id\":\"%s\"}}", utils.DocTypeTradeRequest, request.ShareholderID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var tradeRequests []TradeRequestDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var tradeRequest *TradeRequest
		err = json.Unmarshal(queryResponse.Value, &tradeRequest)
		if err != nil {
			return nil, err
		}
		tradeRequestItem := TradeRequestDoc{Key: queryResponse.Key, Doc: tradeRequest}
		tradeRequests = append(tradeRequests, tradeRequestItem)
	}
	return tradeRequests, nil
}
