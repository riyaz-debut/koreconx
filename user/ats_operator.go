package user

import (
	"bytes"
	"encoding/json"
	"fmt"

	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddATSOperator saves the ATS Operator in world state
func AddATSOperator(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(ATSOperator)

	// change json to strcut
	err := json.Unmarshal([]byte(data), request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeATSOperator
	request.UpdatedAt = request.CreatedAt

	// check the corporate id already exists or not
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"ats_operator_id\":\"%s\"}}", request.DocType, request.ATSOperatorID)
	atsOperatorBA, _, _ := utils.GetByQuery(ctx, queryString, "")

	if atsOperatorBA != nil {
		return nil, status.ErrNotFound.WithMessage(fmt.Sprintf("ATS Operator with ID: %s already exist!", request.ATSOperatorID))
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

// GetATSOperator returns all persons found in world state
func GetATSOperator(ctx contractapi.TransactionContextInterface, data []byte) (*ATSOperator, error) {
	request := new(utils.ResponseID)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	company, err := GetATSOperatorByID(ctx, request.ID)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	return &company, nil
}

// CheckATSOperatorsByID takes an array of IDs,
// checks whether all the IDs are vaild ATS Operators or not
func CheckATSOperatorsByID(ctx contractapi.TransactionContextInterface, IDs []string) (bool, error) {
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
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\": {\"$in\": %s}}}", utils.DocTypeATSOperator, buffer.String())
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return false, status.ErrInternal.WithError(err)
	}

	defer resultsIterator.Close()

	if !resultsIterator.HasNext() {
		return false, status.ErrNotFound.WithMessage("ATS Operator(s) does not exists!")
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
		return false, status.ErrNotFound.WithMessage("ATS Operator(s) does not exists!")
	}

	return true, nil
}

// GetATSOperatorByID returns the ATSOperator with given ID from world state
func GetATSOperatorByID(ctx contractapi.TransactionContextInterface, ID string) (ATSOperator, error) {
	atsOperatorData := ATSOperator{}

	atsOperatorBA, err := ATSOperatorExists(ctx, ID)
	if err != nil {
		return atsOperatorData, err
	}

	// ummarshal the byte array to structure
	err = json.Unmarshal(atsOperatorBA, &atsOperatorData)
	if err != nil {
		return atsOperatorData, status.ErrInternal.WithError(err)
	}

	return atsOperatorData, nil
}

// ATSOperatorExists checks whether the atsOperator with given ID exists or not in world state
func ATSOperatorExists(ctx contractapi.TransactionContextInterface, ID string) ([]byte, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeATSOperator, ID)
	atsOperatorBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("ATSOperator with ID: %s does not exists!", ID))

	return atsOperatorBA, err
}
