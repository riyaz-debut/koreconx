package koresecurities

import (
	"encoding/json"
	"fmt"

	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/user"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddSecuritiesExchangePrice saves the certificate hash in world state
func AddSecuritiesExchangePrice(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(SecuritiesExchangePrice)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeSecuritiesExchangePrice
	request.UpdatedAt = request.CreatedAt

	// company
	_, err = user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// security id
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	var transactionID string

	// Fetch Exchange price
	exchangePrice, dataKey, err := GetExchangePriceBySecuritiesID(ctx, request.KoresecuritiesID)
	if err == nil {
		transactionID = dataKey
		request.CreatedAt = exchangePrice.CreatedAt
	} else {
		transactionID = ctx.GetStub().GetTxID()
	}

	// change strcut to json
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(transactionID, jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseID)
	response.ID = transactionID
	return response, nil
}

// GetExchangePriceBySecuritiesID fetches the exchange prcie for securities with the given ID from world state
func GetExchangePriceBySecuritiesID(ctx contractapi.TransactionContextInterface, ID string) (SecuritiesExchangePrice, string, error) {
	data := SecuritiesExchangePrice{}
	dataBA, dataKey, err := SecuritiesExchangePriceExists(ctx, ID)
	if err != nil {
		return data, "", err
	}

	// ummarshal the byte array to structure
	err = json.Unmarshal(dataBA, &data)
	if err != nil {
		return data, "", status.ErrInternal.WithError(err)
	}
	return data, dataKey, nil
}

// SecuritiesExchangePriceExists check whether the exchange price for securities with given ID exists or not
func SecuritiesExchangePriceExists(ctx contractapi.TransactionContextInterface, ID string) ([]byte, string, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"koresecurities_id\":\"%s\"}}", utils.DocTypeSecuritiesExchangePrice, ID)

	dataBA, dataKey, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("Exchange price for Securities with ID: %s does not exists!", ID))
	return dataBA, dataKey, err
}
