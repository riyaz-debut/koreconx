package koresecurities

import (
	"encoding/json"
	"fmt"

	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/korecontract"
	"kore_chaincode/user"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// IssueSecurities saves the Securities information into the world state
func IssueSecurities(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(IssueSecuritiesInputRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// Check whether company exists or not?

	_, err = user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	// Prepare Securities
	security := Securities{}

	if request.ShareholderAgreementID != "" {
		// Check whether shareholder agreement exists or not?
		_, err = korecontract.GetShareHolderAgreementByID(ctx, request.ShareholderAgreementID, request.CompanyID)
		if err != nil {
			return nil, err
		}
		security.ShareholderAgreementID = request.ShareholderAgreementID
	}

	security.DocType = utils.DocTypeSecurities
	security.CreatedAt = request.CreatedAt
	security.UpdatedAt = request.CreatedAt
	security.Status = request.Status

	// Prepare the SecuritiesCertificate
	securitiesCertificate := SecuritiesCertificate{}
	securitiesCertificate.DateIssued = security.CreatedAt
	securitiesCertificate.CompanyID = request.CompanyID
	securitiesCertificate.SecuritiesholderID = request.CompanyID
	//securitiesCertificate.SecuritiesholderName = company.CompanyLegalName
	securitiesCertificate.CertificateNumber = request.CertificateNumber
	securitiesCertificate.SecuritiesType = request.SecuritiesType
	securitiesCertificate.Status = "1"
	securitiesCertificate.CertificateParent = []string{}
	securitiesCertificate.CertificateSucessor = []string{}

	// Check whether offering memorandum exists or not?
	if request.OfferingMemorandumID != "" {
		offeringMemorandum, err := korecontract.GetOfferingMemorandumByID(ctx, request.OfferingMemorandumID, request.CompanyID)
		if err != nil {
			return nil, err
		}

		security.OfferingMemorandumID = request.OfferingMemorandumID
		securitiesCertificate.ClassOfSecurities = offeringMemorandum.OfferingDetail.OfferingSecuritiesClass
		securitiesCertificate.CompanySymbol = offeringMemorandum.CompanySymbol
		securitiesCertificate.Currency = offeringMemorandum.OfferingDetail.OfferingCurrency
		securitiesCertificate.NumberOfSecurities = offeringMemorandum.OfferingDetail.OfferingSecuritiesNumber
		securitiesCertificate.OfferingPricePerSecurity = offeringMemorandum.OfferingDetail.OfferingPricePerSecurity
		security.AvailableSecurities = securitiesCertificate.NumberOfSecurities
	}

	security.SecuritiesCertificate = securitiesCertificate

	// change strcut to json
	jsonData, err := json.Marshal(security)
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

// GetAllSecurities returns all securities found in world state
func GetAllSecurities(ctx contractapi.TransactionContextInterface, data []byte) ([]SecuritiesDoc, error) {
	request := new(SecuritiesFilter)

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

	queryCompanyID := fmt.Sprintf(",\"securities_certificate.company_id\":\"%s\"", request.CompanyID)
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"%s}}", utils.DocTypeSecurities, queryCompanyID)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var securities []SecuritiesDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var holding *Securities
		err = json.Unmarshal(queryResponse.Value, &holding)
		if err != nil {
			return nil, err
		}
		securitiesItem := SecuritiesDoc{Key: queryResponse.Key, Doc: holding}
		securities = append(securities, securitiesItem)
	}
	return securities, nil
}

// UpdateSecurities saves the holding hash in world state
func UpdateSecurities(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
	request := new(UpdatingSecuritiesRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// securities
	security, err := GetSecuritiesByID(ctx, request.KoresecuritiesID, "")
	if err != nil {
		return nil, err
	}

	security.AvailableSecurities = request.AvailableSecurities
	security.UpdatedAt = request.CreatedAt

	// change strcut to json
	jsonData, err := json.Marshal(security)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.KoresecuritiesID, jsonData)

	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseMessage)
	response.Message = "Offering has been updted successfully!"
	return response, nil
}

// GetSecuritiesByID fetches the Securities with the given ID from world state
func GetSecuritiesByID(ctx contractapi.TransactionContextInterface, ID, CompanyID string) (Securities, error) {
	data := Securities{}
	dataBA, err := SecuritiesExists(ctx, ID, CompanyID)
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

// SecuritiesExists check whether the Securities with given ID exists or not
func SecuritiesExists(ctx contractapi.TransactionContextInterface, ID, CompanyID string) ([]byte, error) {
	var queryString string
	if CompanyID != "" {
		queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\", \"securities_certificate.company_id\": \"%s\"}}", utils.DocTypeSecurities, ID, CompanyID)
	} else {
		queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeSecurities, ID)
	}

	dataBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("Securities with ID: %s does not exists!", ID))
	return dataBA, err
}

// AssociateATSWithSecurity function
func AssociateATSWithSecurity(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
	request := new(AssociateATSWithSecurityRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	company, err := user.GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	_, err = user.GetATSOperatorByID(ctx, request.AtsOperatorID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsATSAssociated(request.AtsOperatorID)
	if !found {
		return nil, err
	}

	// securities
	security, err := GetSecuritiesByID(ctx, request.KoresecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err = security.IsATSAssociated()
	if found {
		return nil, err
	}

	security.AtsOperatorID = request.AtsOperatorID
	security.UpdatedAt = request.CreatedAt

	// change strcut to json
	jsonData, err := json.Marshal(security)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.KoresecuritiesID, jsonData)

	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseMessage)
	response.Message = "Offering has been updted successfully!"
	return response, nil
}

// AssociateBrokerWithSecurity function
func AssociateBrokerWithSecurity(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
	request := new(AssociateBrokerWithSecurityRequest)

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

	_, err = user.GetBrokerDealerByID(ctx, request.BrokerDealerID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsBrokerAssociated(request.BrokerDealerID)
	if !found {
		return nil, err
	}

	// securities
	security, err := GetSecuritiesByID(ctx, request.KoresecuritiesID, request.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err = security.IsBrokerAssociated()
	if found {
		return nil, err
	}

	security.BrokerDealerID = request.BrokerDealerID
	security.UpdatedAt = request.CreatedAt

	// change strcut to json
	jsonData, err := json.Marshal(security)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.KoresecuritiesID, jsonData)

	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseMessage)
	response.Message = "Offering has been updted successfully!"
	return response, nil
}

// IsATSAssociated function
func (data Securities) IsATSAssociated() (bool, error) {
	if data.AtsOperatorID != "" {
		return true, status.ErrInternal.WithMessage("ATS Operator is already associated with the KoreSecurities")
	}
	return false, nil
}

// IsBrokerAssociated function
func (data Securities) IsBrokerAssociated() (bool, error) {
	if data.BrokerDealerID != "" {
		return true, status.ErrInternal.WithMessage("Broker Dealer is already associated with the KoreSecurities")
	}
	return false, nil
}

// GetSecurities returns all persons found in world state
func GetSecurities(ctx contractapi.TransactionContextInterface, data []byte) (*Securities, error) {
	request := new(utils.ResponseID)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	securities, err := GetSecuritiesByID(ctx, request.ID, "")
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	return &securities, nil
}

// GetSecuritiesIDByrequestor function
func GetSecuritiesIDByrequestor(ctx contractapi.TransactionContextInterface, CompanyID, RequestorID string) ([]string, error) {
	queryCompanyID := fmt.Sprintf(",\"securities_certificate.company_id\":\"%s\"", CompanyID)
	queryRequestorID := fmt.Sprintf(",\"$or\":[{\"broker_dealer_id\": \"%s\"},{\"ats_operator_id\": \"%s\"}]", RequestorID, RequestorID)
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"%s%s}}", utils.DocTypeSecurities, queryCompanyID, queryRequestorID)

	securitiesIDs := make([]string, 0)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return securitiesIDs, status.ErrInternal.WithError(err)
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return securitiesIDs, status.ErrInternal.WithError(err)
		}
		securitiesIDs = append(securitiesIDs, queryResponse.Key)
	}
	return securitiesIDs, nil
}
