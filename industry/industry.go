// Package industry defines the industry data structures
package industry

import (
	"encoding/json"
	"fmt"
	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddIndustry saves the Industry information in world state
func AddIndustry(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(Industry)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeIndustry
	request.UpdatedAt = request.CreatedAt

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

// GetAllIndustries returns all industries found in world state
func GetAllIndustries(ctx contractapi.TransactionContextInterface, data []byte) ([]IndustryDoc, error) {
	request := new(RequestAllIndustries)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	var queryIndustry string

	if request.IndustryID != "" {
		queryIndustry = fmt.Sprintf(",\"parent_id\":\"%s\"", request.IndustryID)
	} else {
		queryIndustry = fmt.Sprintf(",\"parent_id\":\"\"")
	}

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"%s}}", utils.DocTypeIndustry, queryIndustry)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var industries []IndustryDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var industry *Industry
		err = json.Unmarshal(queryResponse.Value, &industry)
		if err != nil {
			return nil, err
		}
		industryItem := IndustryDoc{Key: queryResponse.Key, Doc: industry}
		industries = append(industries, industryItem)
	}
	return industries, nil
}
