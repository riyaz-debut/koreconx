package user

import (
	"bytes"
	"encoding/json"
	"fmt"

	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddTransferAgent saves the transfer agent in world state
func AddTransferAgent(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(TransferAgent)

	// change json to strcut
	err := json.Unmarshal([]byte(data), request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeTransferAgent
	request.UpdatedAt = request.CreatedAt

	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"transfer_agent_id\":\"%s\"}}", request.DocType, request.TransferAgentID)
	transferAgentBA, _, _ := utils.GetByQuery(ctx, queryString, "")
	if transferAgentBA != nil {
		return nil, status.ErrNotFound.WithMessage(fmt.Sprintf("Transfer Agent with ID: %s already exist!", request.TransferAgentID))
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

// CheckTransferAgentByID takes input as an array of IDs,
// checks whether all IDs are valid transfer agent or not
func CheckTransferAgentByID(ctx contractapi.TransactionContextInterface, IDs []string) (bool, error) {
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

	// Check the Transfer Agent list whether it is correct or not
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\": {\"$in\": %s}}}", utils.DocTypeTransferAgent, buffer.String())
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return false, status.ErrInternal.WithError(err)
	}

	defer resultsIterator.Close()

	if !resultsIterator.HasNext() {
		return false, status.ErrNotFound.WithMessage("Transfer Agent(s) does not exists!")
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
		return false, status.ErrNotFound.WithMessage("Transfer Agent(s) does not exists!")
	}
	return true, nil
}

// GetTransferAgentByID returns the TransferAgent with given ID from world state
func GetTransferAgentByID(ctx contractapi.TransactionContextInterface, ID string) (TransferAgent, error) {
	transferAgentData := TransferAgent{}

	transferAgentBA, err := TransferAgentExists(ctx, ID)
	if err != nil {
		return transferAgentData, err
	}

	// ummarshal the byte array to structure
	err = json.Unmarshal(transferAgentBA, &transferAgentData)
	if err != nil {
		return transferAgentData, status.ErrInternal.WithError(err)
	}

	return transferAgentData, nil
}

// TransferAgentExists checks whether the transferAgent with given ID exists or not in world state
func TransferAgentExists(ctx contractapi.TransactionContextInterface, ID string) ([]byte, error) {
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeTransferAgent, ID)
	transferAgentBA, _, err := utils.GetByQuery(ctx, queryString, fmt.Sprintf("TransferAgent with ID: %s does not exists!", ID))

	return transferAgentBA, err
}

// GetTransferAgent returns all persons found in world state
func GetTransferAgent(ctx contractapi.TransactionContextInterface, data []byte) (*TransferAgent, error) {
	request := new(utils.ResponseID)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	company, err := GetTransferAgentByID(ctx, request.ID)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	return &company, nil
}
