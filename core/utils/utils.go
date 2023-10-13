package utils

import (
	"fmt"
	"kore_chaincode/core/status"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// MetaData common fields which are used in all other structures
type MetaData struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DocType   string    `json:"doc_type"`
}

// ResponseID is used to return the response which contains only one ID field
type ResponseID struct {
	ID string `json:"id"`
}

// ResponseIDArray data fields
type ResponseIDArray struct {
	Data []string `json:"data"`
}

// ResponseMessage is used to return the response which contains only one message field
type ResponseMessage struct {
	Message string `json:"message"`
}

// GetByQuery executes the query and returns the byte array result
func GetByQuery(ctx contractapi.TransactionContextInterface, query string, message string) ([]byte, string, error) {
	fmt.Println("<===== GetByQuery =====>")
	fmt.Println("Query is:", query)
	fmt.Println("Message is:", message)

	stub := ctx.GetStub()

	resultsIterator, err := stub.GetQueryResult(query)
	if err != nil {
		return nil, "", status.ErrInternal.WithError(err)
	}

	defer resultsIterator.Close()

	if !resultsIterator.HasNext() {
		fmt.Println("Record does not exists in DB")
		return nil, "", status.ErrNotFound.WithMessage(message)
	}

	queryResponse, err := resultsIterator.Next()
	if err != nil {
		return nil, "", status.ErrInternal.WithError(err)
	}

	fmt.Println("Query Response:", queryResponse)
	return queryResponse.Value, queryResponse.Key, nil
}
