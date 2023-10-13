// Package user defines the user data structures
package user

import (
	"bytes"
	"encoding/json"
	"fmt"

	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddServiceProvider saves the ServiceProvider information into the world state
func AddServiceProvider(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(ServiceProvider)

	// change json to strcut
	err := json.Unmarshal([]byte(data), request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeServiceProvider
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

// GetAllServiceProviders returns all service providers found in world state
func GetAllServiceProviders(ctx contractapi.TransactionContextInterface, data []byte) ([]ServiceProviderDoc, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"}}", utils.DocTypeServiceProvider)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var compaines []ServiceProviderDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var serviceProvider *ServiceProvider
		err = json.Unmarshal(queryResponse.Value, &serviceProvider)
		if err != nil {
			return nil, err
		}
		serviceProviders := ServiceProviderDoc{Key: queryResponse.Key, Doc: serviceProvider}
		compaines = append(compaines, serviceProviders)
	}

	return compaines, nil
}

// GetServiceProvider returns all persons found in world state
func GetServiceProvider(ctx contractapi.TransactionContextInterface, data []byte) (*ServiceProvider, error) {
	request := new(utils.ResponseID)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	company, err := GetServiceProviderByID(ctx, request.ID)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	return &company, nil
}

// GetServiceProviderByID returns the serviceProvider with given ID from world state
func GetServiceProviderByID(ctx contractapi.TransactionContextInterface, ID string) (ServiceProvider, error) {
	serviceProviderData := ServiceProvider{}

	serviceProviderBA, err := ServiceProviderExists(ctx, ID)
	if err != nil {
		return serviceProviderData, err
	}

	// ummarshal the byte array to structure
	err = json.Unmarshal(serviceProviderBA, &serviceProviderData)
	if err != nil {
		return serviceProviderData, status.ErrInternal.WithError(err)
	}

	return serviceProviderData, nil
}

// ServiceProviderExists checks whether the company with given ID exists or not in world state
func ServiceProviderExists(ctx contractapi.TransactionContextInterface, ID string) ([]byte, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeServiceProvider, ID)
	companyBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("Service Provider with ID: %s does not exists!", ID))
	return companyBA, err
}

// CheckServiceProvidersByID takes an array of IDs,
// checks whether all the IDs are vaild ATS Operators or not
func CheckServiceProvidersByID(ctx contractapi.TransactionContextInterface, IDs []string) (bool, error) {
	var buffer bytes.Buffer
	totalLength := len(IDs)

	buffer.WriteString("[")
	for i := 0; i < totalLength; i++ {
		buffer.WriteString("\"")
		buffer.WriteString(IDs[i])
		buffer.WriteString("\"")
		if i != totalLength-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")

	// Check the ATS Operators list whether it is correct or not
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\": {\"$in\": %s}}}", utils.DocTypeServiceProvider, buffer.String())
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return false, status.ErrInternal.WithError(err)
	}

	defer resultsIterator.Close()

	if !resultsIterator.HasNext() {
		return false, status.ErrNotFound.WithMessage("Service Provider(s) does not exists!")
	}

	fetchedRecords := 0

	for resultsIterator.HasNext() {
		_, err := resultsIterator.Next()
		if err != nil {
			return false, status.ErrInternal.WithError(err)
		}
		fetchedRecords++
	}

	// check for total results
	if fetchedRecords != totalLength {
		return false, status.ErrNotFound.WithMessage("Service Provider(s) does not exists!")
	}

	return true, nil
}
