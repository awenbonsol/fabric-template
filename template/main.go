package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		log.Panicf("Error creating chaincode: %v", err)
	}

	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
