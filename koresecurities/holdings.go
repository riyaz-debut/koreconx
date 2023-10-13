// Package koresecurities Holding data
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

// AddHolding saves the holding hash in world state
func AddHolding(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(TransactionRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// company
	_, err = user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// securities holder id
	_, err = person.GetPersonByID(ctx, request.SecuritiesHolderID)
	if err != nil {
		return nil, err
	}

	// securitiy id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// Create the transaction
	transaction := Transaction{}
	// set the default values for the fields
	transaction.CompanyID = request.CompanyID
	transaction.SecuritiesHolderID = request.SecuritiesHolderID
	transaction.KoresecuritiesID = request.KoresecuritiesID
	transaction.CertificateNumber = request.CertificateNumber
	transaction.HoldingAmount = request.HoldingAmount
	transaction.AveragePrice = request.AveragePrice
	transaction.DateAcquired = request.DateAcquired
	transaction.SourceSystemID = request.SourceSystemID
	transaction.CreatedAt = request.CreatedAt
	transaction.DocType = utils.DocTypeTransaction
	transaction.UpdatedAt = request.CreatedAt

	// change strcut to json
	jsonData, err := json.Marshal(transaction)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(ctx.GetStub().GetTxID(), jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// Holding
	holding, holdingID, err := GetHoldingBySecuritiesID(ctx, transaction.SecuritiesHolderID, transaction.KoresecuritiesID)

	if holdingID != "" {
		holding.HoldingAmount = holding.HoldingAmount + transaction.HoldingAmount
		holding.AvailableShares = holding.AvailableShares + transaction.HoldingAmount
		holding.UpdatedAt = transaction.CreatedAt
	} else {
		holding = Holding{}
		holding.CompanyID = transaction.CompanyID
		holding.SecuritiesHolderID = transaction.SecuritiesHolderID
		holding.KoresecuritiesID = transaction.KoresecuritiesID
		holding.HoldingAmount = transaction.HoldingAmount
		holding.AvailableShares = transaction.HoldingAmount
		holding.CreatedAt = transaction.CreatedAt
		holding.DocType = utils.DocTypeHolding
		holding.LastUpdatedAt = transaction.CreatedAt
		holding.UpdatedAt = transaction.CreatedAt
		holding.ReasonCode = ""
		holdingID = request.TransactionID
	}

	// change strcut to json
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

// GetNumberOfSharesInHolding returns all holdings found in world state
func GetNumberOfSharesInHolding(ctx contractapi.TransactionContextInterface, data []byte) ([]HoldingDoc, error) {
	request := new(HoldingFilter)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// company
	company, err := user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// securities holder id
	_, err = person.GetPersonByID(ctx, request.SecuritiesHolderID)
	if err != nil {
		return nil, err
	}

	// securitiy id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsRequestorAssociated(request.RequestorID)
	if !found {
		return nil, err
	}

	var queryString string
	queryCompanyID, querySecuritiesID, querySecuritiesHolderID := "", "", ""
	queryCompanyID = fmt.Sprintf(",\"company_id\":\"%s\"", request.CompanyID)
	querySecuritiesID = fmt.Sprintf(",\"koresecurities_id\":\"%s\"", request.KoresecuritiesID)
	querySecuritiesHolderID = fmt.Sprintf(",\"securities_holder_id\":\"%s\"", request.SecuritiesHolderID)

	queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"%s%s%s}}", utils.DocTypeHolding, queryCompanyID, querySecuritiesID, querySecuritiesHolderID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var holdings []HoldingDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var holding *Holding
		err = json.Unmarshal(queryResponse.Value, &holding)
		if err != nil {
			return nil, err
		}

		holdingItem := HoldingDoc{HoldingID: queryResponse.Key, HoldingAmount: holding.HoldingAmount, AvailableShares: holding.AvailableShares}
		holdings = append(holdings, holdingItem)
	}

	return holdings, nil
}

// GetAllHoldingsbyAllSecurities returns all holdings found in world state
func GetAllHoldingsbyAllSecurities(ctx contractapi.TransactionContextInterface, data []byte) ([]AllHoldingAllSecuritiesResponse, error) {
	request := new(AllHoldingsByAllSecuritiesRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// company
	company, err := user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsRequestorAssociated(request.RequestorID)
	if !found {
		return nil, err
	}

	var queryString string
	queryCompanyID, querySecuritiesID := "", ""
	queryCompanyID = fmt.Sprintf(",\"company_id\":\"%s\"", request.CompanyID)

	isTA := company.IsRequestorTA(request.RequestorID)
	if !isTA {
		securitiesID, err := GetSecuritiesIDByrequestor(ctx, request.CompanyID, request.RequestorID)
		if err != nil {
			return nil, err
		}

		securitiesIDsJSON, err := json.Marshal(securitiesID)
		if err != nil {
			return nil, status.ErrInternal.WithError(err)
		}

		querySecuritiesID = fmt.Sprintf(",\"koresecurities_id\": {\"$in\": %s}", securitiesIDsJSON)
	}

	queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"%s%s}}", utils.DocTypeHolding, queryCompanyID, querySecuritiesID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var holdings []AllHoldingAllSecuritiesResponse
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var holding *Holding
		err = json.Unmarshal(queryResponse.Value, &holding)
		if err != nil {
			return nil, err
		}

		holdingItem := AllHoldingAllSecuritiesResponse{}
		holdingItem.KoreSecuritiesID = holding.KoresecuritiesID
		holdingItem.HoldingID = queryResponse.Key
		holdingItem.SecuritiesHolderID = holding.SecuritiesHolderID
		holdings = append(holdings, holdingItem)
	}

	return holdings, nil
}

// GetAllHoldingsbySecuritiesID returns all holdings found in world state
func GetAllHoldingsbySecuritiesID(ctx contractapi.TransactionContextInterface, data []byte) ([]AllHoldingBySecuritiesResponse, error) {
	request := new(GetAllHoldingsbySecuritiesIDRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// company
	company, err := user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsRequestorAssociated(request.RequestorID)
	if !found {
		return nil, err
	}

	var queryString string
	queryCompanyID, querySecuritiesID := "", ""
	queryCompanyID = fmt.Sprintf(",\"company_id\":\"%s\"", request.CompanyID)
	querySecuritiesID = fmt.Sprintf(",\"koresecurities_id\":\"%s\"", request.KoreSecuritiesID)

	queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"%s%s}}", utils.DocTypeHolding, queryCompanyID, querySecuritiesID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var holdings []AllHoldingBySecuritiesResponse
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var holding *Holding
		err = json.Unmarshal(queryResponse.Value, &holding)
		if err != nil {
			return nil, err
		}

		holdingItem := AllHoldingBySecuritiesResponse{}
		holdingItem.HoldingID = queryResponse.Key
		holdingItem.SecuritiesHolderID = holding.SecuritiesHolderID
		holdings = append(holdings, holdingItem)
	}

	return holdings, nil
}

// GetAllShareHolders returns all holdings found in world state
func GetAllShareHolders(ctx contractapi.TransactionContextInterface, data []byte) (*ShareHoldersIDResponse, error) {
	request := new(ShareholderFilter)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// company
	company, err := user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// securitiy id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsRequestorAssociated(request.RequestorID)
	if !found {
		return nil, err
	}

	var queryString string

	queryCompanyID, querySecuritiesID := "", ""
	queryCompanyID = fmt.Sprintf(",\"company_id\":\"%s\"", request.CompanyID)
	querySecuritiesID = fmt.Sprintf(",\"koresecurities_id\":\"%s\"", request.KoresecuritiesID)

	queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"%s%s}}", utils.DocTypeHolding, queryCompanyID, querySecuritiesID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	shareHoldersID := []string{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var holding *Holding
		err = json.Unmarshal(queryResponse.Value, &holding)
		if err != nil {
			return nil, err
		}

		shareHoldersID = append(shareHoldersID, holding.SecuritiesHolderID)
	}

	response := ShareHoldersIDResponse{Data: shareHoldersID}

	return &response, nil
}

// GetAllShareHoldersByComapny returns all holdings found in world state
func GetAllShareHoldersByComapny(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseIDArray, error) {
	request := new(AllHoldingsByAllSecuritiesRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// company
	company, err := user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsRequestorAssociated(request.RequestorID)
	if !found {
		return nil, err
	}

	var queryString string
	queryCompanyID, querySecuritiesID := "", ""
	queryCompanyID = fmt.Sprintf(",\"company_id\":\"%s\"", request.CompanyID)

	isTA := company.IsRequestorTA(request.RequestorID)
	if !isTA {
		securitiesID, err := GetSecuritiesIDByrequestor(ctx, request.CompanyID, request.RequestorID)
		if err != nil {
			return nil, err
		}

		securitiesIDsJSON, err := json.Marshal(securitiesID)
		if err != nil {
			return nil, status.ErrInternal.WithError(err)
		}

		querySecuritiesID = fmt.Sprintf(",\"koresecurities_id\": {\"$in\": %s}", securitiesIDsJSON)
	}

	queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"%s%s}}", utils.DocTypeHolding, queryCompanyID, querySecuritiesID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// f2 := make([]string, 0)
	var shareHolders []string
	var shareHoldersMap = make(map[string]bool)
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var holding *Holding
		err = json.Unmarshal(queryResponse.Value, &holding)
		if err != nil {
			return nil, err
		}

		// to create unique slice
		if !shareHoldersMap[holding.SecuritiesHolderID] {
			shareHolders = append(shareHolders, holding.SecuritiesHolderID)
			shareHoldersMap[holding.SecuritiesHolderID] = true
		}
	}

	response := utils.ResponseIDArray{Data: shareHolders}
	return &response, nil
}

// GetAvailableShares saves the holding hash in world state
func GetAvailableShares(ctx contractapi.TransactionContextInterface, data []byte) (*ResponseInvestorHoldingExists, error) {
	request := new(GetAvailableSharesRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// securitiy id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, "")
	if err != nil {
		return nil, err
	}

	// securities holder id
	_, err = person.GetPersonByID(ctx, request.SecuritiesHolderID)
	if err != nil {
		return nil, err
	}

	holding, _, err := GetHoldingBySecuritiesID(ctx, request.SecuritiesHolderID, request.KoresecuritiesID)
	if err != nil {
		return nil, err
	}

	company, err := user.GetCompanyByID(ctx, holding.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsRequestorAssociated(request.RequestorID)
	if !found {
		return nil, err
	}

	response := new(ResponseInvestorHoldingExists)
	response.Exists = false

	if holding.AvailableShares >= request.NumberOfShares {
		response.Exists = true
	}

	return response, nil
}

// PlaceHoldOnShares saves the holding hash in world state
func PlaceHoldOnShares(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(HoldSharesTransaction)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// securitiy id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, "")
	if err != nil {
		return nil, err
	}

	// securities holder id
	_, err = person.GetPersonByID(ctx, request.SecuritiesHolderID)
	if err != nil {
		return nil, err
	}

	// set default values
	request.DocType = utils.DocTypeHoldShares
	request.Status = "0"
	request.AvailableNumberOfShares = request.NumberOfShares
	request.UpdatedAt = request.CreatedAt

	holding, holdingID, err := GetHoldingBySecuritiesID(ctx, request.SecuritiesHolderID, request.KoresecuritiesID)
	if err != nil {
		return nil, err
	}

	if holding.AvailableShares < request.NumberOfShares {
		return nil, status.ErrInternal.WithMessage("Shareholder does not have enough shares in his holding!")
	}

	holding.AvailableShares = holding.AvailableShares - request.NumberOfShares
	holding.UpdatedAt = request.CreatedAt
	holding.LastUpdatedAt = request.LastUpdatedAt
	// holding.ReasonCode = request.ReasonCode

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
	jsonData, err = json.Marshal(request)
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

// ReleaseHoldOnShares saves the holding hash in world state
func ReleaseHoldOnShares(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(ReleaseSharesTransaction)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// securitiy id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, "")
	if err != nil {
		return nil, err
	}

	// securities holder id
	_, err = person.GetPersonByID(ctx, request.SecuritiesHolderID)
	if err != nil {
		return nil, err
	}

	// set default values
	request.DocType = utils.DocTypeReleaseShares
	request.UpdatedAt = request.CreatedAt

	holding, holdingID, err := GetHoldingBySecuritiesID(ctx, request.SecuritiesHolderID, request.KoresecuritiesID)
	if err != nil {
		return nil, err
	}

	// hold share transaction id
	holdShares, _, err := GetHoldSharesByID(ctx, request.ATSTransactionID, request.SecuritiesHolderID, request.KoresecuritiesID)
	if err != nil {
		return nil, err
	}

	if holdShares.AvailableNumberOfShares < request.NumberOfShares {
		return nil, status.ErrInternal.WithMessage("Number of hold shares are less than the number of shares to be released.")
	}

	holding.AvailableShares = holding.AvailableShares + request.NumberOfShares
	holding.LastUpdatedAt = request.LastUpdatedAt
	holding.UpdatedAt = request.CreatedAt
	// holding.ReasonCode = request.ReasonCode

	// save holding
	jsonData, err := json.Marshal(holding)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(holdingID, jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// save release transaction
	jsonData, err = json.Marshal(request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(ctx.GetStub().GetTxID(), jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// update the hold share request
	holdShares.AvailableNumberOfShares = holdShares.AvailableNumberOfShares - request.NumberOfShares

	if holdShares.AvailableNumberOfShares == 0 {
		holdShares.Status = "1"
	}

	jsonData, err = json.Marshal(holdShares)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.ATSTransactionID, jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseID)
	response.ID = ctx.GetStub().GetTxID()
	return response, nil
}

// UpdateHolding saves the holding hash in world state
func UpdateHolding(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(UpdateHoldingRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// securitiy id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, "")
	if err != nil {
		return nil, err
	}

	// securities holder id
	_, err = person.GetPersonByID(ctx, request.SecuritiesHolderID)
	if err != nil {
		return nil, err
	}

	holding, holdingID, err := GetHoldingBySecuritiesID(ctx, request.SecuritiesHolderID, request.KoresecuritiesID)
	if err != nil {
		return nil, err
	}

	holding.HoldingAmount = request.NumberOfShares
	holding.AvailableShares = request.NumberOfShares
	holding.LastUpdatedAt = request.LastUpdatedAt
	holding.UpdatedAt = request.CreatedAt
	holding.ReasonCode = request.ReasonCode

	// change strcut to json
	jsonData, err := json.Marshal(holding)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(holdingID, jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseID)
	response.ID = holdingID
	return response, nil
}

// InvestorHoldingExists checks if the investor has a particular holding or not
func InvestorHoldingExists(ctx contractapi.TransactionContextInterface, data []byte) (*ResponseInvestorHoldingExists, error) {
	request := new(InvestorHoldingRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// securitiy id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, "")
	if err != nil {
		return nil, err
	}

	// securities holder id
	_, err = person.GetPersonByID(ctx, request.SecuritiesHolderID)
	if err != nil {
		return nil, err
	}

	holding, _, err := GetHoldingBySecuritiesID(ctx, request.SecuritiesHolderID, request.KoresecuritiesID)
	if err != nil {
		return nil, err
	}

	company, err := user.GetCompanyByID(ctx, holding.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsRequestorAssociated(request.RequestorID)
	if !found {
		return nil, err
	}

	response := new(ResponseInvestorHoldingExists)
	response.Exists = true
	return response, nil
}

// GetHoldingBySecuritiesID returns the holdings of a shareholder with given ID from world state
func GetHoldingBySecuritiesID(ctx contractapi.TransactionContextInterface, SecuritiesHolderID, KoresecuritiesID string) (Holding, string, error) {
	holdingData := Holding{}

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"securities_holder_id\":\"%s\",\"koresecurities_id\":\"%s\"}}", utils.DocTypeHolding, SecuritiesHolderID, KoresecuritiesID)

	holdingBA, holdingID, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("The investor does not hold this koresecurities!"))

	// ummarshal the byte array to structure
	err = json.Unmarshal(holdingBA, &holdingData)
	if err != nil {
		return holdingData, "", status.ErrInternal.WithError(err)
	}

	return holdingData, holdingID, nil
}

// GetHoldSharesByID returns the holdings of a shareholder with given ID from world state
func GetHoldSharesByID(ctx contractapi.TransactionContextInterface, HoldShareID, SecuritiesHolderID, KoresecuritiesID string) (HoldSharesTransaction, string, error) {
	holdingData := HoldSharesTransaction{}

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"securities_holder_id\":\"%s\",\"koresecurities_id\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeHoldShares, SecuritiesHolderID, KoresecuritiesID, HoldShareID)

	holdingBA, holdingID, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("The hold share transaction does not exists!"))

	if err != nil {
		return holdingData, "", err
	}

	// ummarshal the byte array to structure
	err = json.Unmarshal(holdingBA, &holdingData)
	if err != nil {
		return holdingData, "", status.ErrInternal.WithError(err)
	}

	return holdingData, holdingID, nil
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// GetAllHoldingsByCompany returns all holdings found in world state
func GetAllHoldingsByCompany(ctx contractapi.TransactionContextInterface, data []byte) ([]HoldingByCompanyResponse, error) {
	request := new(ShareholderFilter)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// company
	company, err := user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// securitiy id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsRequestorAssociated(request.RequestorID)
	if !found {
		return nil, err
	}

	var queryString string

	queryCompanyID, querySecuritiesID := "", ""
	queryCompanyID = fmt.Sprintf(",\"company_id\":\"%s\"", request.CompanyID)
	querySecuritiesID = fmt.Sprintf(",\"koresecurities_id\":\"%s\"", request.KoresecuritiesID)

	queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"%s%s}}", utils.DocTypeHolding, queryCompanyID, querySecuritiesID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var holdings []HoldingByCompanyResponse
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var holding *Holding
		err = json.Unmarshal(queryResponse.Value, &holding)
		if err != nil {
			return nil, err
		}

		holdingItem := HoldingByCompanyResponse{HoldingAmount: holding.HoldingAmount, SecuritiesHolderID: holding.SecuritiesHolderID, AvailableShares: holding.AvailableShares}
		holdings = append(holdings, holdingItem)
	}

	return holdings, nil
}

// GetAllTradableHoldings returns all holdings found in world state
func GetAllTradableHoldings(ctx contractapi.TransactionContextInterface, data []byte) ([]AllTradableHoldingsResponse, error) {
	request := new(AllTradableHoldingsRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// fetch all the companies with which the requestor is associated
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"ats_operators\":{\"$elemMatch\":{\"$eq\":\"%s\"}}}}", utils.DocTypeCompany, request.RequestorID)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var companyIDs []string
	var companies = make(map[string]user.Company)
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var company *user.Company
		err = json.Unmarshal(queryResponse.Value, &company)
		if err != nil {
			return nil, err
		}

		companyIDs = append(companyIDs, queryResponse.Key)
		companies[queryResponse.Key] = *company
	}

	if len(companyIDs) == 0 {
		return nil, status.ErrBadRequest.WithMessage("Requestor does not have any associated company!")
	}

	// fetch the all the holdings of the shareholder and company
	companyIDsJSON, err := json.Marshal(companyIDs)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"securities_holder_id\":\"%s\",\"company_id\": {\"$in\": %s}}}", utils.DocTypeHolding, request.SecuritiesHolderID, companyIDsJSON)

	resultsIterator, err = ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var response []AllTradableHoldingsResponse
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var holding *Holding
		err = json.Unmarshal(queryResponse.Value, &holding)
		if err != nil {
			return nil, err
		}

		// securitiy id
		security, err := GetSecuritiesByID(ctx, holding.KoresecuritiesID, holding.CompanyID)
		if err != nil {
			return nil, err
		}

		//company := companies[holding.CompanyID]
		holdingItem := AllTradableHoldingsResponse{}
		//oldingItem.CompanyLegalName = company.CompanyLegalName
		holdingItem.OfferingSecuritiesClass = security.SecuritiesCertificate.ClassOfSecurities
		holdingItem.CompanySymbol = security.SecuritiesCertificate.CompanySymbol
		holdingItem.HoldingAmount = holding.HoldingAmount
		holdingItem.AvailableShares = holding.AvailableShares
		holdingItem.SecuritiesHolderID = holding.SecuritiesHolderID
		holdingItem.KoresecuritiesID = holding.KoresecuritiesID
		holdingItem.HoldingID = queryResponse.Key
		response = append(response, holdingItem)
	}
	return response, nil
}
