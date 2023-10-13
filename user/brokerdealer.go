package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddBrokerDealer saves the Broker Dealer in world state
func AddBrokerDealer(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(BrokerDealer)

	// change json to strcut
	err := json.Unmarshal([]byte(data), request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeBrokerDealer
	request.UpdatedAt = request.CreatedAt

	// check the company id already exists or not
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"broker_dealer_id\":\"%s\"}}", request.DocType, request.BrokerDealerID)
	brokerDealerBA, _, _ := utils.GetByQuery(ctx, queryString, "")
	if brokerDealerBA != nil {
		return nil, status.ErrNotFound.WithMessage(fmt.Sprintf("Broker Dealer with ID: %s already exist!", request.BrokerDealerID))
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

// GetBrokerDealer returns all persons found in world state
func GetBrokerDealer(ctx contractapi.TransactionContextInterface, data []byte) (*BrokerDealer, error) {
	request := new(utils.ResponseID)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	company, err := GetBrokerDealerByID(ctx, request.ID)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	return &company, nil
}

// GetBrokerDealerByID fetches the Broker Dealer with given ID from world state
func GetBrokerDealerByID(ctx contractapi.TransactionContextInterface, ID string) (BrokerDealer, error) {
	brokerDealerData := BrokerDealer{}

	brokerDealerBA, err := BrokerDealerExists(ctx, ID)
	if err != nil {
		return brokerDealerData, err
	}

	err = json.Unmarshal(brokerDealerBA, &brokerDealerData)
	if err != nil {
		return brokerDealerData, status.ErrInternal.WithError(err)
	}

	return brokerDealerData, nil
}

// BrokerDealerExists checks whether the Broker dealer with the given ID exists or not
func BrokerDealerExists(ctx contractapi.TransactionContextInterface, ID string) ([]byte, error) {
	query := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\":\"%s\"}}", utils.DocTypeBrokerDealer, ID)
	brokerDealerBA, _, err := utils.GetByQuery(ctx, query, fmt.Sprintf("Broker Dealer with ID: %s does not exists!", ID))
	return brokerDealerBA, err
}

// CheckBrokerDealerByID takes input an array of IDs,
// check whether all IDs are valid Broker dealer or not
func CheckBrokerDealerByID(ctx contractapi.TransactionContextInterface, IDs []string) (bool, error) {
	// make the string for the ats operator id
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

	// Check the Broker dealer list whether it is correct or not
	queryString := fmt.Sprintf("{\"selector\":{\"doc_type\":\"%s\",\"_id\": {\"$in\": %s}}}", utils.DocTypeBrokerDealer, buffer.String())
	resultsIterator, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil {
		return false, status.ErrInternal.WithError(err)
	}
	defer resultsIterator.Close()

	if !resultsIterator.HasNext() {
		return false, status.ErrNotFound.WithMessage("Broker Dealer(s) does not exists!")
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
		return false, status.ErrNotFound.WithMessage("Broker Dealer(s) does not exists!")
	}

	return true, nil
}
