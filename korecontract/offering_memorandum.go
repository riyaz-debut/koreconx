package korecontract

import (
	"encoding/json"
	"fmt"
	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/user"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddOfferingMemorandum saves the Offering Memorandum in world state
func AddOfferingMemorandum(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(OfferingMemorandum)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeOfferingMemorandum
	request.UpdatedAt = request.CreatedAt

	// company
	_, err = user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// document
	_, err = GetDocumentByID(ctx, request.DocumentID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// broker dealers

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

// GetAllOfferingMemorandums returns all offeringMemorandums found in world state
func GetAllOfferingMemorandums(ctx contractapi.TransactionContextInterface, data []byte) ([]OfferingMemorandumDoc, error) {
	request := new(CompanyFilter)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	_, err = user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"company_id\":\"%s\"}}", utils.DocTypeOfferingMemorandum, request.CompanyID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var offeringMemorandums []OfferingMemorandumDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var offeringMemorandum *OfferingMemorandum
		err = json.Unmarshal(queryResponse.Value, &offeringMemorandum)
		if err != nil {
			return nil, err
		}
		offeringMemorandumItem := OfferingMemorandumDoc{Key: queryResponse.Key, Doc: offeringMemorandum}
		offeringMemorandums = append(offeringMemorandums, offeringMemorandumItem)
	}
	return offeringMemorandums, nil
}

// GetOfferingMemorandumByID fetches the offeringMemorandum with the given ID from world state
func GetOfferingMemorandumByID(ctx contractapi.TransactionContextInterface, ID, CompanyID string) (OfferingMemorandum, error) {
	data := OfferingMemorandum{}
	dataBA, err := OfferingMemorandumExists(ctx, ID, CompanyID)
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

// OfferingMemorandumExists checks whether the offeringMemorandum with the given ID exists or not
func OfferingMemorandumExists(ctx contractapi.TransactionContextInterface, ID, CompanyID string) ([]byte, error) {
	var queryString string

	if CompanyID != "" {
		queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\", \"company_id\": \"%s\"}}", utils.DocTypeOfferingMemorandum, ID, CompanyID)
	} else {
		queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeOfferingMemorandum, ID)
	}

	dataBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("OfferingMemorandum with ID: %s does not exists!", ID))
	return dataBA, err
}
