package korecontract

import (
	"encoding/json"
	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"
	"kore_chaincode/person"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddPaymentMethod saves the document hash in world state
func AddPaymentMethod(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(PaymentMethod)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypePaymentMethod
	request.UpdatedAt = request.CreatedAt

	// investor id
	_, err = person.GetPersonByID(ctx, request.PayerID)
	if err != nil {
		return nil, err
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
