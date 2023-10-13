// Package user defines the user data structures
package user

import (
	"encoding/json"
	"fmt"

	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddCompany saves the company information in world state
func AddCompany(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	company := new(Company)

	// change json to strcut
	err := json.Unmarshal([]byte(data), company)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	company.DocType = utils.DocTypeCompany
	company.UpdatedAt = company.CreatedAt

	// // check the company already exists or not
	// queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"company_legal_name\":\"%s\",\"legal_extension\":\"%s\",\"country\":\"%s\"}}", utils.DocTypeCompany, company.CompanyLegalName, company.LegalExtension, company.Country)
	// companyBA, _, _ := utils.GetByQuery(ctx, queryString, "")

	// if companyBA != nil {
	// 	return nil, status.ErrStatusConflict.WithMessage(fmt.Sprintf("Company with company_legal_name: %s, legal_extension %s & country %s already exists!", company.CompanyLegalName, company.LegalExtension, company.Country))
	// }

	// Set default values
	company.ATSOperators = []string{}
	//company.NotificationURLs = []string{}
	company.BrokerDealers = []string{}
	company.ServiceProviders = []string{}

	// change strcut to json
	jsonData, err := json.Marshal(company)
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

// ImportCompanies adds a new person in world state
func ImportCompanies(ctx contractapi.TransactionContextInterface, data []byte) ([]CompanyWithID, error) {
	companies := new(ImportCompaniesRequest)

	// change json to strcut
	err := json.Unmarshal(data, companies)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := []CompanyWithID{}

	for i := 0; i < len(companies.Data); i++ {
		// set the default values for the fields
		companies.Data[i].Data.DocType = utils.DocTypeCompany
		companies.Data[i].Data.UpdatedAt = companies.Data[i].Data.CreatedAt

		// // check the company already exists or not
		// queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"company_legal_name\":\"%s\",\"legal_extension\":\"%s\",\"country\":\"%s\"}}", utils.DocTypeCompany, companies.Data[i].Data.CompanyLegalName, companies.Data[i].Data.LegalExtension, companies.Data[i].Data.Country)
		// companyBA, _, _ := utils.GetByQuery(ctx, queryString, "")

		// if companyBA != nil {
		// 	continue
		// }

		// Set default values
		companies.Data[i].Data.ATSOperators = []string{}
		// companies.Data[i].Data.NotificationURLs = []string{}
		companies.Data[i].Data.BrokerDealers = []string{}
		companies.Data[i].Data.ServiceProviders = []string{}

		// change strcut to json
		jsonData, err := json.Marshal(companies.Data[i].Data)
		if err != nil {
			return nil, status.ErrInternal.WithError(err)
		}

		err = ctx.GetStub().PutState(companies.Data[i].ID, jsonData)

		if err != nil {
			return nil, status.ErrInternal.WithError(err)
		}

		newCompanyID := CompanyWithID{ID: companies.Data[i].ID, Data: companies.Data[i].Data}
		response = append(response, newCompanyID)
	}
	return response, nil
}

// GetAllCompanies returns all persons found in world state
func GetAllCompanies(ctx contractapi.TransactionContextInterface, data []byte) ([]CompanyDoc, error) {
	companyFilter := new(CompanyFilter)

	// change json to strcut
	err := json.Unmarshal(data, companyFilter)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	queryCountry, queryBrokerDealerID, queryIndustryID, querySubIndustryID := "", "", "", ""
	queryCompanyName := ""

	// if companyFilter.Country != "" {
	// 	queryCountry = fmt.Sprintf(",\"country\":\"%s\"", companyFilter.Country)
	// }

	if companyFilter.BrokerDealerID != "" {
		queryBrokerDealerID = fmt.Sprintf(",\"broker_dealers\":{\"$elemMatch\":{\"$eq\":\"%s\"}}", companyFilter.BrokerDealerID)
	}

	// if companyFilter.IndustryID != "" {
	// 	queryCountry = fmt.Sprintf(",\"industry_id\":\"%s\"", companyFilter.IndustryID)
	// }

	// if companyFilter.SubIndustryID != "" {
	// 	queryCountry = fmt.Sprintf(",\"sub_industry_id\":\"%s\"", companyFilter.SubIndustryID)
	// }

	// if companyFilter.CompanyLegalName != "" {
	// 	queryCompanyName = fmt.Sprintf(",\"company_legal_name\":\"%s\"", companyFilter.CompanyLegalName)
	// }

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"%s%s%s%s%s}}", utils.DocTypeCompany, queryCountry, queryBrokerDealerID, queryIndustryID, querySubIndustryID, queryCompanyName)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var compaines []CompanyDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var company *Company
		err = json.Unmarshal(queryResponse.Value, &company)
		if err != nil {
			return nil, err
		}
		listResponse := CompanyDoc{Key: queryResponse.Key, Doc: company}
		compaines = append(compaines, listResponse)
	}

	return compaines, nil
}

// GetAllCompaniesByRequestorID returns all companies found in world state
func GetAllCompaniesByRequestorID(ctx contractapi.TransactionContextInterface, data []byte) ([]CompanyList, error) {
	companyFilter := new(CompaniesByRequestorID)

	// change json to strcut
	err := json.Unmarshal(data, companyFilter)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	queryAtsOperator := fmt.Sprintf("{\"ats_operators\":{\"$elemMatch\":{\"$eq\":\"%s\"}}}", companyFilter.RequestorID)
	queryBrokerDealer := fmt.Sprintf("{\"broker_dealers\":{\"$elemMatch\":{\"$eq\":\"%s\"}}}", companyFilter.RequestorID)
	queryTransferAgent := fmt.Sprintf("{\"transfer_agent_id\":\"%s\"}", companyFilter.RequestorID)
	queryServiceProvider := fmt.Sprintf("{\"service_providers\":{\"$elemMatch\":{\"$eq\":\"%s\"}}}", companyFilter.RequestorID)

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"$or\":[%s,%s,%s,%s]}}", utils.DocTypeCompany, queryAtsOperator, queryBrokerDealer, queryTransferAgent, queryServiceProvider)
	fmt.Println("queryString", queryString)
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var list []CompanyList
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var company *Company
		err = json.Unmarshal(queryResponse.Value, &company)
		if err != nil {
			return nil, err
		}

		responseData := CompanyList{}
		responseData.ID = queryResponse.Key
		//responseData.CompanyLegalName = company.CompanyLegalName
		list = append(list, responseData)
	}

	return list, nil
}

// AssociateNotificationURLWithCompany returns all persons found in world state
func AssociateNotificationURLWithCompany(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
	request := new(AssociateNotificationURLRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	company, err := GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// notificationURLs := company.NotificationURLs
	// notificationURLs = append(notificationURLs, request.URL)
	// company.NotificationURLs = notificationURLs

	// change strcut to json
	jsonData, err := json.Marshal(company)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.CompanyID, jsonData)

	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseMessage)
	response.Message = "URL has been associated with the comapnu!"
	return response, nil
}

// GetCompany returns all persons found in world state
func GetCompany(ctx contractapi.TransactionContextInterface, data []byte) (*Company, error) {
	request := new(utils.ResponseID)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	company, err := GetCompanyByID(ctx, request.ID)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	return &company, nil
}

// UpdateCompany updates the company information in world state
func UpdateCompany(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(UpdateCompanyRequest)

	// change json to strcut
	err := json.Unmarshal([]byte(data), request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	companyID := request.ID

	// check if the company exists or not
	company, err := GetCompanyByID(ctx, companyID)
	if err != nil {
		return nil, err
	}

	// industry id and sub industry id

	// check the company already exists or not
	// queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"company_legal_name\":\"%s\",\"legal_extension\":\"%s\",\"country\":\"%s\",\"_id\":{\"$ne\":\"%s\"}}}", utils.DocTypeCompany, request.Data.CompanyLegalName, request.Data.LegalExtension, request.Data.Country, companyID)
	// companyBA, _, _ := utils.GetByQuery(ctx, queryString, "")

	// if companyBA != nil {
	// 	return nil, status.ErrStatusConflict.WithMessage(fmt.Sprintf("Company with company_legal_name: %s, legal_extension %s & country %s already exists!", request.Data.CompanyLegalName, request.Data.LegalExtension, request.Data.Country))
	// }

	// set the default values for the fields
	request.Data.UpdatedAt = request.Data.CreatedAt
	request.Data.CreatedAt = company.CreatedAt
	request.Data.DocType = company.DocType
	// request.Data.BankruptcyProceeding = company.BankruptcyProceeding
	// request.Data.RegulatoryInjunction = company.RegulatoryInjunction
	// request.Data.HoldByATSOperator = company.HoldByATSOperator
	request.Data.ATSOperators = company.ATSOperators
	request.Data.TransferAgentID = company.TransferAgentID
	request.Data.BrokerDealers = company.BrokerDealers
	// request.Data.NotificationURLs = company.NotificationURLs

	// change strcut to json
	jsonData, err := json.Marshal(request.Data)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.ID, jsonData)

	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseID)
	response.ID = request.ID
	return response, nil
}

// CompanyBankruptcyProceedingStatus updates the Company Bankruptcy Proceeding status in Korechian
// func CompanyBankruptcyProceedingStatus(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
// 	return updateCompanyStatus(ctx, data, 1)
// }

// CompanyRegulatoryInjunctionStatus updates the Company Regulatory Injunction status in world state
// func CompanyRegulatoryInjunctionStatus(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
// 	return updateCompanyStatus(ctx, data, 2)
// }

// CompanyHoldByATSOperatorStatus updates the Company Hold By ATS Operator status in Korechian
// func CompanyHoldByATSOperatorStatus(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
// 	return updateCompanyStatus(ctx, data, 3)
// }

// AssociateATSOperator Associates the ATS Operators with the Company
func AssociateATSOperator(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
	return associateOperatorsWithCompany(ctx, data, 1)
}

// AssociateBrokerDealer Associates the Brokerdealer with the Company
func AssociateBrokerDealer(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
	return associateOperatorsWithCompany(ctx, data, 2)
}

// AssociateServiceProvider Associates the ATS Operators with the Company
func AssociateServiceProvider(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
	return associateOperatorsWithCompany(ctx, data, 3)
}

// AssociateTransferAgent Associates the TransferAgent with the Company
func AssociateTransferAgent(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
	request := new(AssociateTransferAgentRequest)

	// change json to strcut
	err := json.Unmarshal([]byte(data), request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// check if the company exists or not
	company, err := GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	found, err := company.IsTAAssociated()
	if found {
		return nil, err
	}

	_, err = GetTransferAgentByID(ctx, request.TransferAgentID)
	if err != nil {
		return nil, err
	}

	company.TransferAgentID = request.TransferAgentID

	// change strcut to json
	jsonData, err := json.Marshal(company)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.CompanyID, jsonData)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseMessage)
	response.Message = "Company Information updated successfully!"
	return response, nil
}

// AssignManagementPeople assigns the director to company
// func AssignManagementPeople(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
// 	request := new(CompanyManagementRequest)

// 	// change json to strcut
// 	err := json.Unmarshal([]byte(data), request)
// 	if err != nil {
// 		return nil, status.ErrInternal.WithError(err)
// 	}

// 	// check the company already exists or not
// 	companyData, err := GetCompanyByID(ctx, request.CompanyID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	management := companyData.Management
// 	if management == nil {
// 		management = &Management{}
// 	}

// 	var details map[string]ManagementPeople
// 	people := ManagementPeople{PersonID: request.PersonID, ShareholderVoteTransactionID: request.ShareholderVoteTransactionID, Status: "1", StartDate: request.StartDate}

// 	switch request.Role {
// 	case "ceo":
// 		if len(management.CEO) == 0 {
// 			details = make(map[string]ManagementPeople)
// 		} else {
// 			details = management.CEO
// 		}
// 		details[request.PersonID] = people
// 		management.CEO = details
// 	case "cfo":
// 		if len(management.CFO) == 0 {
// 			details = make(map[string]ManagementPeople)
// 		} else {
// 			details = management.CFO
// 		}
// 		details[request.PersonID] = people
// 		management.CFO = details
// 	case "director":
// 		if len(management.Director) == 0 {
// 			details = make(map[string]ManagementPeople)
// 		} else {
// 			details = management.Director
// 		}
// 		details[request.PersonID] = people
// 		management.Director = details
// 	case "officer":
// 		if len(management.Officer) == 0 {
// 			details = make(map[string]ManagementPeople)
// 		} else {
// 			details = management.Officer
// 		}
// 		details[request.PersonID] = people
// 		management.Officer = details
// 	default:
// 		err = status.ErrInternal.WithMessage("Entered role is wrong")
// 	}

// 	if err != nil {
// 		return nil, err
// 	}

// 	companyData.Management = management
// 	companyData.UpdatedAt = request.CreatedAt

// 	// change strcut to json
// 	jsonData, err := json.Marshal(companyData)
// 	if err != nil {
// 		return nil, status.ErrInternal.WithError(err)
// 	}

// 	err = ctx.GetStub().PutState(request.CompanyID, jsonData)

// 	if err != nil {
// 		return nil, status.ErrInternal.WithError(err)
// 	}

// 	response := new(utils.ResponseMessage)
// 	response.Message = "Company Information updated successfully!"
// 	return response, nil
// }

// RemoveManagementPeople removes the director to company
// func RemoveManagementPeople(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseMessage, error) {
// 	request := new(CompanyManagementRemoveRequest)

// 	// change json to strcut
// 	err := json.Unmarshal([]byte(data), request)
// 	if err != nil {
// 		return nil, status.ErrInternal.WithError(err)
// 	}

// 	// check the company already exists or not
// 	companyData, err := GetCompanyByID(ctx, request.CompanyID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var details map[string]ManagementPeople

// 	management := companyData.Management

// 	if management == nil {
// 		err = status.ErrInternal.WithMessage("Management is not updated.")
// 		return nil, err
// 	}

// 	switch request.Role {
// 	case "ceo":
// 		details = companyData.Management.CEO
// 		val, ok := details[request.PersonID]

// 		if !ok {
// 			err = status.ErrInternal.WithMessage("Person is not in the Management list")
// 			return nil, err
// 		}

// 		val.Status = "0"
// 		val.EndDate = request.EndDate
// 		details[request.PersonID] = val
// 		companyData.Management.CEO = details
// 	case "cfo":
// 		details = companyData.Management.CEO
// 		val, ok := details[request.PersonID]

// 		if !ok {
// 			err = status.ErrInternal.WithMessage("Person is not in the Management list")
// 			return nil, err
// 		}

// 		val.Status = "0"
// 		val.EndDate = request.EndDate
// 		details[request.PersonID] = val
// 		companyData.Management.CFO = details
// 	case "director":
// 		details = companyData.Management.CEO
// 		val, ok := details[request.PersonID]

// 		if !ok {
// 			err = status.ErrInternal.WithMessage("Person is not in the Management list")
// 			return nil, err
// 		}

// 		val.Status = "0"
// 		val.EndDate = request.EndDate
// 		details[request.PersonID] = val
// 		companyData.Management.Director = details
// 	case "officer":
// 		details = companyData.Management.CEO
// 		val, ok := details[request.PersonID]

// 		if !ok {
// 			err = status.ErrInternal.WithMessage("Person is not in the Management list")
// 			return nil, err
// 		}

// 		val.Status = "0"
// 		val.EndDate = request.EndDate
// 		details[request.PersonID] = val
// 		companyData.Management.Officer = details
// 	default:
// 		err = status.ErrInternal.WithMessage("Entered role is wrong")
// 		return nil, err
// 	}

// 	companyData.UpdatedAt = request.CreatedAt

// 	// change strcut to json
// 	jsonData, err := json.Marshal(companyData)
// 	if err != nil {
// 		return nil, status.ErrInternal.WithError(err)
// 	}

// 	err = ctx.GetStub().PutState(request.CompanyID, jsonData)

// 	if err != nil {
// 		return nil, status.ErrInternal.WithError(err)
// 	}

// 	response := new(utils.ResponseMessage)
// 	response.Message = "Company Information updated successfully!"
// 	return response, nil
// }

// GetManagementPeople fetch management people
// func GetManagementPeople(ctx contractapi.TransactionContextInterface, data []byte) (*Management, error) {
// 	request := new(GetManagementPeopleRequest)

// 	// change json to strcut
// 	err := json.Unmarshal([]byte(data), request)
// 	if err != nil {
// 		return nil, status.ErrInternal.WithError(err)
// 	}

// 	// check the company already exists or not
// 	companyData, err := GetCompanyByID(ctx, request.CompanyID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return companyData.Management, nil
// }

// GetCompanyByID returns the company with given ID from world state
func GetCompanyByID(ctx contractapi.TransactionContextInterface, ID string) (Company, error) {
	companyData := Company{}

	companyBA, err := CompanyExists(ctx, ID)
	if err != nil {
		return companyData, err
	}

	// ummarshal the byte array to structure
	err = json.Unmarshal(companyBA, &companyData)
	if err != nil {
		return companyData, status.ErrInternal.WithError(err)
	}

	// if len(companyData.NotificationURLs) == 0 {
	// 	companyData.NotificationURLs = []string{}
	// }

	return companyData, nil
}

// CompanyExists checks whether the company with given ID exists or not in world state
func CompanyExists(ctx contractapi.TransactionContextInterface, ID string) ([]byte, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeCompany, ID)
	companyBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("Company with ID: %s does not exists!", ID))

	return companyBA, err
}

// updateCompanyStatus
// func updateCompanyStatus(ctx contractapi.TransactionContextInterface, data []byte, actionType int) (*utils.ResponseMessage, error) {
// 	request := new(CompanyStatusRequest)

// 	// change json to strcut
// 	err := json.Unmarshal([]byte(data), request)
// 	if err != nil {
// 		return nil, status.ErrInternal.WithError(err)
// 	}

// 	// check the company already exists or not
// 	companyData, err := GetCompanyByID(ctx, request.CompanyID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	details := make(map[string]ReferenceDetails)
// 	for _, referenceDetail := range request.ReferenceDetails {
// 		details[referenceDetail.AuthorizingEntityID] = referenceDetail
// 	}

// 	// create the object to store in company profile
// 	issuerStatus := BankruptcyProceeding{}
// 	issuerStatus.Status = request.Status
// 	issuerStatus.ReferenceDetails = details

// 	// update the data
// 	switch actionType {
// 	case 1:
// 		companyData.BankruptcyProceeding = &issuerStatus
// 	case 2:
// 		companyData.RegulatoryInjunction = &issuerStatus
// 	case 3:
// 		companyData.HoldByATSOperator = &issuerStatus

// 	}
// 	companyData.UpdatedAt = request.CreatedAt

// 	// change strcut to json
// 	jsonData, err := json.Marshal(companyData)
// 	if err != nil {
// 		return nil, status.ErrInternal.WithError(err)
// 	}

// 	err = ctx.GetStub().PutState(request.CompanyID, jsonData)

// 	if err != nil {
// 		return nil, status.ErrInternal.WithError(err)
// 	}

// 	response := new(utils.ResponseMessage)
// 	response.Message = "Company Information updated successfully!"
// 	return response, nil
// }

// associateOperatorsWithCompany
func associateOperatorsWithCompany(ctx contractapi.TransactionContextInterface, data []byte, actionType int) (*utils.ResponseMessage, error) {
	request := new(AssociateIDsRequest)

	// change json to strcut
	err := json.Unmarshal([]byte(data), request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// check if the company exists or not
	company, err := GetCompanyByID(ctx, request.CompanyID)
	if err != nil {
		return nil, err
	}

	switch actionType {
	case 1:
		exists, err := CheckATSOperatorsByID(ctx, request.Data)
		if !exists {
			return nil, err
		}
		company.ATSOperators = request.Data
	case 2:
		exists, err := CheckBrokerDealerByID(ctx, request.Data)
		if !exists {
			return nil, err
		}
		company.BrokerDealers = request.Data
	case 3:
		exists, err := CheckServiceProvidersByID(ctx, request.Data)
		if !exists {
			return nil, err
		}
		company.ServiceProviders = request.Data
	}

	// change strcut to json
	jsonData, err := json.Marshal(company)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	err = ctx.GetStub().PutState(request.CompanyID, jsonData)

	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := new(utils.ResponseMessage)
	response.Message = "Company Information updated successfully!"
	return response, nil
}

// InBankruptcyProceeding returns whether the company is in Bankruptcy Proceeding or not
// func (data Company) InBankruptcyProceeding() (bool, error) {
// 	if data.BankruptcyProceeding != nil {
// 		details := *data.BankruptcyProceeding
// 		if details.Status {
// 			return true, status.ErrBadRequest.WithMessage("Company is in bankruptcy proceedings.")
// 		}
// 	}
// 	return false, nil
// }

// InRegulatoryInjunction returns whether the company is in Regulatory Injunction or not
// func (data Company) InRegulatoryInjunction() (bool, error) {
// 	if data.RegulatoryInjunction != nil {
// 		details := *data.RegulatoryInjunction
// 		if details.Status {
// 			return true, status.ErrBadRequest.WithMessage("Company is in regulatory injunction.")
// 		}
// 	}
// 	return false, nil
// }

// InHoldByATSOperator returns whether the company is in hold by ATS Operator or not
// func (data Company) InHoldByATSOperator() (bool, error) {
// 	if data.HoldByATSOperator != nil {
// 		details := *data.HoldByATSOperator
// 		if details.Status {
// 			return true, status.ErrBadRequest.WithMessage("Company is in hold by ATS Operator.")
// 		}
// 	}
// 	return false, nil
// }

// InGoodStandingWithTransferAgent returns whether the company is in good standing with trasnfer agent
func (data Company) InGoodStandingWithTransferAgent() (bool, error) {
	if data.TransferAgentID != "" {
		return true, nil
	}
	return false, status.ErrBadRequest.WithMessage("Company is not in good standing with the transfer agent.")
}

// Status returns whether the comapny is allowed to transact or not
// func (data Company) Status() (bool, error) {
// 	// // InBankruptcyProceeding
// 	// if ok, err := data.InBankruptcyProceeding(); ok {
// 	// 	return false, err
// 	// }

// 	// InRegulatoryInjunction
// 	// if ok, err := data.InRegulatoryInjunction(); ok {
// 	// 	return false, err
// 	// }

// 	// HoldByATSOperator
// 	if ok, err := data.InHoldByATSOperator(); ok {
// 		return false, err
// 	}
// 	return true, nil
// }

// IsATSAssociated checks whether the passed ATS ID is associated with company or not
func (data Company) IsATSAssociated(AtsOperatorID string) (bool, error) {
	_, found := Find(data.ATSOperators, AtsOperatorID)
	if !found {
		return false, status.ErrInternal.WithMessage("ATS Operator is not associated with the company")
	}
	return true, nil

}

// IsBrokerAssociated checks whether the passed Broker ID is associated with company or not
func (data Company) IsBrokerAssociated(BrokerDealerID string) (bool, error) {
	_, found := Find(data.BrokerDealers, BrokerDealerID)
	if !found {
		return false, status.ErrInternal.WithMessage("Broker Dealer is not associated with the company")
	}
	return true, nil

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

// IsTAAssociated function
func (data Company) IsTAAssociated() (bool, error) {
	if data.TransferAgentID != "" {
		return true, status.ErrInternal.WithMessage("Transfer Agent is already associated with the Company")
	}

	return false, nil
}

// IsRequestorAssociated function
func (data Company) IsRequestorAssociated(ID string) (bool, error) {
	taFound := data.TransferAgentID == ID
	baFound, _ := data.IsBrokerAssociated(ID)
	atsFound, _ := data.IsATSAssociated(ID)

	if !taFound && !baFound && !atsFound {
		return false, status.ErrInternal.WithMessage("Requestor is not associated with the company")
	}

	return true, nil
}

// IsRequestorTA function
func (data Company) IsRequestorTA(ID string) bool {
	return data.TransferAgentID == ID
}
