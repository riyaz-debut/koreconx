package person

import (
	"encoding/json"
	"fmt"
	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/user"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddPerson adds a new person in world state
func AddPerson(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	person := new(Person)

	// change json to strcut
	err := json.Unmarshal(data, person)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	person.DocType = utils.DocTypePerson
	person.UpdatedAt = person.CreatedAt

	if person.Verifications.KYCVerification != nil {
		_, err = user.GetServiceProviderByID(ctx, person.Verifications.KYCVerification.ProviderID)
		if err != nil {
			return nil, err
		}
	}

	// change strcut to json
	jsonData, err := json.Marshal(person)
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

// ImportPerson adds a new person in world state
func ImportPerson(ctx contractapi.TransactionContextInterface, data []byte) ([]PersonWithID, error) {
	persons := new(ImportPersonRequest)

	// change json to strcut
	err := json.Unmarshal(data, persons)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	response := []PersonWithID{}

	for i := 0; i < len(persons.Data); i++ {
		// set the default values for the fields
		persons.Data[i].Data.DocType = utils.DocTypePerson
		persons.Data[i].Data.UpdatedAt = persons.Data[i].Data.CreatedAt

		// change strcut to json
		jsonData, err := json.Marshal(persons.Data[i].Data)
		if err != nil {
			return nil, status.ErrInternal.WithError(err)
		}

		err = ctx.GetStub().PutState(persons.Data[i].ID, jsonData)

		if err != nil {
			return nil, status.ErrInternal.WithError(err)
		}

		newPerson := PersonWithID{ID: persons.Data[i].ID, Data: persons.Data[i].Data}
		response = append(response, newPerson)
	}
	return response, nil
}

// GetAllPersons returns all persons found in world state
func GetAllPersons(ctx contractapi.TransactionContextInterface) ([]PersonDoc, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\"}}", utils.DocTypePerson)

	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var persons []PersonDoc
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var person *Person
		err = json.Unmarshal(queryResponse.Value, &person)
		if err != nil {
			return nil, err
		}
		personItem := PersonDoc{Key: queryResponse.Key, Doc: person}
		persons = append(persons, personItem)
	}
	return persons, nil
}

// GetPerson returns all persons found in world state
func GetPerson(ctx contractapi.TransactionContextInterface, data []byte) (*Person, error) {
	request := new(ShowPersonRequest)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	person, err := GetPersonByID(ctx, request.ID)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	return &person, nil
}

// GetPersonByID fetches the document with the given ID from world state
func GetPersonByID(ctx contractapi.TransactionContextInterface, ID string) (Person, error) {
	data := Person{}
	dataBA, err := ExistsByID(ctx, ID)
	if err != nil {
		return data, err
	}

	// ummarshal the byte array to structure
	err = json.Unmarshal(dataBA, &data)
	if err != nil {
		return data, status.ErrInternal.WithError(err)
	}
	fmt.Println("Person", data)
	return data, nil
}

// ExistsByID checks whether the document with the given ID exists or not
func ExistsByID(ctx contractapi.TransactionContextInterface, ID string) ([]byte, error) {
	fmt.Println("<===== ExistsByID =====>")
	fmt.Println("ID is:", ID)

	var queryString string

	queryString = fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypePerson, ID)

	dataBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("Person with ID: %s does not exists!", ID))
	return dataBA, err
}

// HasVerifications bool
func (person Person) HasVerifications() bool {
	return person.Verifications != nil
}

// IsKYCVerification bool
func (person Person) IsKYCVerification(date time.Time) (bool, error) {
	errorMessage := status.ErrBadRequest.WithMessage("Person does not have the KYC Verification.")

	if !(person.HasVerifications() && person.HasKYCVerification()) {
		return false, errorMessage
	}

	if date.IsZero() {
		return true, nil
	}

	verficiationDetails := person.Verifications.KYCVerification
	if verficiationDetails.VerificationDate.Before(date) && verficiationDetails.VerificationExpiryDate.After(date) {
		return true, nil
	}

	return false, errorMessage
}

// HasKYCVerification bool
func (person Person) HasKYCVerification() bool {
	return person.Verifications.KYCVerification != nil
}
