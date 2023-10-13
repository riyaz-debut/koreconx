package korecontract

import (
	"encoding/json"
	"fmt"
	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddDocument saves the document hash in world state
func AddDocument(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(Document)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeDocument
	request.UpdatedAt = request.CreatedAt

	// // company
	// _, err = user.GetCompanyByID(ctx, request.EntityID)
	// if err != nil {
	// 	return nil, err
	// }

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

// GetAllDocuments returns all documents found in world state
func GetAllDocuments(ctx contractapi.TransactionContextInterface, data []byte) ([]DocumentDoc, error) {
	request := new(DocumentFilter)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// _, err = user.GetCompanyByID(ctx, request.EntityID)
	// if err != nil {
	// 	return nil, err
	// }

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"entity_id\":\"%s\"}}", utils.DocTypeDocument, request.EntityID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var documents []DocumentDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var document *Document
		err = json.Unmarshal(queryResponse.Value, &document)
		if err != nil {
			return nil, err
		}
		documentItem := DocumentDoc{Key: queryResponse.Key, Doc: document}
		documents = append(documents, documentItem)
	}
	return documents, nil
}

// GetDocumentByID fetches the document with the given ID from world state
func GetDocumentByID(ctx contractapi.TransactionContextInterface, ID, EntityID string) (Document, error) {
	data := Document{}
	dataBA, err := DocumentExists(ctx, ID, EntityID)
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

// DocumentExists checks whether the document with the given ID exists or not
func DocumentExists(ctx contractapi.TransactionContextInterface, ID, EntityID string) ([]byte, error) {
	var queryString string

	if EntityID != "" {
		queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\", \"entity_id\": \"%s\"}}", utils.DocTypeDocument, ID, EntityID)
	} else {
		queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeDocument, ID)
	}

	dataBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("Document with ID: %s does not exists!", ID))
	return dataBA, err
}
