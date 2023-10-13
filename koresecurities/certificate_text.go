// Package koresecurities securities related information
package koresecurities

import (
	"encoding/json"
	"fmt"

	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/user"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddSecuritiesCertificateText saves the securitiesCertificateText hash in world state
func AddSecuritiesCertificateText(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(SecuritiesCertificateText)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeSecuritiesCertificateText
	request.UpdatedAt = request.CreatedAt

	// company
	_, err = user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// securities
	_, err = GetSecuritiesByID(ctx, request.KoreSecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
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

// GetAllSecuritiesCertificateTexts returns all securitiesCertificateTexts found in world state
func GetAllSecuritiesCertificateTexts(ctx contractapi.TransactionContextInterface, data []byte) ([]SecuritiesCertificateTextDoc, error) {
	request := new(CertificateTextFilter)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"koresecurities_id\":\"%s\"}}", utils.DocTypeSecuritiesCertificateText, request.KoreSecuritiesID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var securitiesCertificateTexts []SecuritiesCertificateTextDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var securitiesCertificateText *SecuritiesCertificateText
		err = json.Unmarshal(queryResponse.Value, &securitiesCertificateText)
		if err != nil {
			return nil, err
		}
		securitiesCertificateTextItem := SecuritiesCertificateTextDoc{Key: queryResponse.Key, Doc: securitiesCertificateText}
		securitiesCertificateTexts = append(securitiesCertificateTexts, securitiesCertificateTextItem)
	}
	return securitiesCertificateTexts, nil
}
