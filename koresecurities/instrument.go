package koresecurities

import (
	"encoding/json"

	"kore_chaincode/core/status"
	"kore_chaincode/core/utils"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// AddSecuritiesInstrument saves the document hash in world state
func AddSecuritiesInstrument(ctx contractapi.TransactionContextInterface, data []byte) (*utils.ResponseID, error) {
	request := new(SecuritiesInstrument)

	// change json to strcut
	err := json.Unmarshal(data, request)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}

	// set the default values for the fields
	request.DocType = utils.DocTypeSecuritiesInstrument
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
