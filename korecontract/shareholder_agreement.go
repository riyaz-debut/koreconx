package korecontract

import (
	"encoding/json"
	"fmt"
	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/user"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddShareHolderAgreement saves the shareHolderAgreement hash in world state
func AddShareHolderAgreement(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(ShareHolderAgreement)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeShareHolderAgreement
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

// GetAllShareHolderAgreements returns all shareHolderAgreements found in world state
func GetAllShareHolderAgreements(ctx contractapi.TransactionContextInterface, data []byte) ([]ShareHolderAgreementDoc, error) {
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

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"company_id\":\"%s\"}}", utils.DocTypeShareHolderAgreement, request.CompanyID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var shareHolderAgreements []ShareHolderAgreementDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var shareHolderAgreement *ShareHolderAgreement
		err = json.Unmarshal(queryResponse.Value, &shareHolderAgreement)
		if err != nil {
			return nil, err
		}
		shareHolderAgreementItem := ShareHolderAgreementDoc{Key: queryResponse.Key, Doc: shareHolderAgreement}
		shareHolderAgreements = append(shareHolderAgreements, shareHolderAgreementItem)
	}
	return shareHolderAgreements, nil
}

// GetShareHolderAgreementByID fetches the shareholder agreement with the given ID from world state
func GetShareHolderAgreementByID(ctx contractapi.TransactionContextInterface, ID, CompanyID string) (ShareHolderAgreement, error) {
	data := ShareHolderAgreement{}
	dataBA, err := ShareHolderAgreementExists(ctx, ID, CompanyID)
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

// ShareHolderAgreementExists checks whether the Shareholder agreement exists or not
func ShareHolderAgreementExists(ctx contractapi.TransactionContextInterface, ID, CompanyID string) ([]byte, error) {
	var queryString string
	if CompanyID != "" {
		queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\", \"company_id\": \"%s\"}}", utils.DocTypeShareHolderAgreement, ID, CompanyID)
	} else {
		queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeShareHolderAgreement, ID)
	}

	dataBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("ShareHolderAgreement with ID: %s does not exists!", ID))
	return dataBA, err
}
