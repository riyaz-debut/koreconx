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

// AddCertificate saves the certificate hash in world state
func AddCertificate(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(Certificate)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeCertificate
	request.UpdatedAt = request.CreatedAt

	// company
	_, err = user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// securities
	_, err = GetSecuritiesByID(ctx, request.KoresecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// securities holder
	_, err = person.GetPersonByID(ctx, request.SecuritiesHolderID)
	if err != nil {
		return nil, err
	}

	// koretransaction

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

// GetAllCertificates returns all certificates found in world state
func GetAllCertificates(ctx contractapi.TransactionContextInterface, data []byte) ([]CertificateDoc, error) {
	request := new(HoldingFilter)

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

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"company_id\":\"%s\"}}", utils.DocTypeCertificate, request.CompanyID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var certificates []CertificateDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var certificate *Certificate
		err = json.Unmarshal(queryResponse.Value, &certificate)
		if err != nil {
			return nil, err
		}
		certificateItem := CertificateDoc{Key: queryResponse.Key, Doc: certificate}
		certificates = append(certificates, certificateItem)
	}
	return certificates, nil
}

// UpdateCertificate updates the Securities certificate in the world state
func UpdateCertificate(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(UpdateCertificateRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// certificte
	certificate, err := GetCertificateByID(ctx, request.CertificateID)
	if err != nil {
		return nil, err
	}

	certificate.Status = request.Status
	certificate.UpdatedAt = request.CreatedAt

	// change strcut to json
	jsonData, err := json.Marshal(certificate)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.CertificateID, jsonData)

	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseID)
	response.ID = request.CertificateID
	return response, nil
}

// GetCertificateByID fetches the Securities certificates with the given ID from world state
func GetCertificateByID(ctx contractapi.TransactionContextInterface, ID string) (Certificate, error) {
	data := Certificate{}
	dataBA, err := CertificateExists(ctx, ID)
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

// CertificateExists checks whether the Securities certificate exists or not
func CertificateExists(ctx contractapi.TransactionContextInterface, ID string) ([]byte, error) {
	var queryString string
	queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeCertificate, ID)

	dataBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("Certificate with ID: %s does not exists!", ID))
	return dataBA, err
}
