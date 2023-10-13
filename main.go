package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// KoreChainCode koreconx chaincode implementation
type KoreChainCode struct {
	contractapi.Contract
}

func main() {
	KoreChainCode, err := contractapi.NewChaincode(&KoreChainCode{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	if err := KoreChainCode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
}
