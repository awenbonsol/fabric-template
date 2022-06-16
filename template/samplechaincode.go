package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
)

type SmartContract struct {
	contractapi.Contract
}

type Person struct {
	FirstName           string                `json:"firstName"`
	LastName            string                `json:"lastName"`
	Age                 int                   `json:"age"`
	NationalID          string                `json:"nationalId"`
	SocialMediaAccounts []SocialMediaAccounts `json:"socialMediaAccounts"`
	IsMarried           bool                  `json:"isMarried"`
}

type SocialMediaAccounts struct {
	WebsiteUrl  string `json:"websiteUrl"`
	DateCreated string `json:"dateCreated"`
}

type QueryResult struct {
	Key    string `json:"Key"`
	Record *Person
}

const (
	firstNameIndex            = "firstName"
	lastNameAndIsMarriedIndex = "lastName~isMarried"
	ageIndex                  = "age"
	zero                      = 0
	emptyString               = ""
)

// CREATE PERSON
func (s *SmartContract) CreatePerson(ctx contractapi.TransactionContextInterface,
	person Person) error {

	logger := logging.NewLogger("samplechaincode")
	logger.Infoln("Start: Calling CreatePerson function.")

	if len(person.NationalID) > 12 {
		return fmt.Errorf("national id %s must not exceed 12 characters", person.NationalID)
	}

	key := person.NationalID
	isExist, err := s.IsExists(ctx, key)
	if err != nil {
		return err
	}

	if isExist {
		return fmt.Errorf("CreatePerson: the person %s is already existing. ", key)
	}

	person = Person{
		FirstName:           person.FirstName,
		LastName:            person.LastName,
		Age:                 person.Age,
		NationalID:          person.NationalID,
		SocialMediaAccounts: person.SocialMediaAccounts,
		IsMarried:           person.IsMarried,
	}

	personAsBytes, err := json.Marshal(person)
	if err != nil {
		return fmt.Errorf("CreatePerson: unable to Marshal %s ", personAsBytes)
	}

	err = ctx.GetStub().PutState(key, personAsBytes)
	if err != nil {
		return fmt.Errorf("CreatePerson: unable to invoke person to ledger")
	}

	firstNameIndexKey, _ := ctx.GetStub().CreateCompositeKey(firstNameIndex, []string{strings.ToLower(person.FirstName), key})
	lastNameAndIsMarriedIndexKey, _ := ctx.GetStub().CreateCompositeKey(strings.ToLower(lastNameAndIsMarriedIndex),
		[]string{person.LastName, strconv.FormatBool(person.IsMarried), key})

	value := []byte{0x00}
	err = ctx.GetStub().PutState(firstNameIndexKey, value)
	if err != nil {
		return fmt.Errorf("CreatePerson: unable to invoke firstNameIndexKey to ledger")
	}

	return ctx.GetStub().PutState(lastNameAndIsMarriedIndexKey, value)
}

// UPDATE PERSON
func (s *SmartContract) UpdatePerson(ctx contractapi.TransactionContextInterface,
	nationalId string, lastName string) error {

	logger := logging.NewLogger("samplechaincode")
	logger.Infoln("Start: Calling UpdatePerson function.")

	queryResult, err := s.GetByNationalId(ctx, nationalId)
	if err != nil {
		return err
	}

	queryResult.LastName = lastName
	queryResult.IsMarried = true

	personAsBytes, err := json.Marshal(queryResult)
	if err != nil {
		return fmt.Errorf("UpdatePerson: unable to Marshal %s ", personAsBytes)
	}

	err = ctx.GetStub().PutState(nationalId, personAsBytes)
	if err != nil {
		return fmt.Errorf("UpdatePerson: unable to invoke person to ledger")
	}

	lastNameAndIsMarriedIndexKey, _ := ctx.GetStub().CreateCompositeKey(lastNameAndIsMarriedIndex,
		[]string{strings.ToLower(queryResult.LastName), strconv.FormatBool(queryResult.IsMarried), nationalId})

	value := []byte{0x00}
	return ctx.GetStub().PutState(lastNameAndIsMarriedIndexKey, value)
}

// GET BY NATIONAL ID
func (s *SmartContract) GetByNationalId(ctx contractapi.TransactionContextInterface,
	nationalId string) (*Person, error) {

	logger := logging.NewLogger("samplechaincode")
	logger.Infoln("Start: Calling GetByNationalId function.")

	if nationalId == emptyString {
		return nil, fmt.Errorf("GetByNationalId: input parameters must not be empty")
	}

	queryResult, err := ctx.GetStub().GetState(nationalId)
	if err != nil {
		return nil, fmt.Errorf("GetByNationalId: failed to read from world state: %v", err)
	}

	if queryResult == nil {
		return nil, fmt.Errorf("GetByNationalId: the person %s does not exist", nationalId)
	}

	var person Person
	err = json.Unmarshal(queryResult, &person)
	if err != nil {
		return nil, err
	}
	logger.Infoln("End: GetByNationalId called with key value of: ", nationalId)
	return &person, nil
}

// GET BY FIRST NAME
func (s *SmartContract) GetByFirstName(ctx contractapi.TransactionContextInterface,
	firstName string) ([]Person, error) {

	logger := logging.NewLogger("samplechaincode")
	logger.Infoln("Start: Calling GetByFirstName function.")

	if firstName == emptyString {
		return nil, fmt.Errorf("GetByFirstName: input parameter must not be empty")
	}

	queryResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(firstNameIndex, []string{strings.ToLower(firstName)})
	if err != nil {
		return nil, err
	}
	defer queryResultsIterator.Close()

	var people []Person
	for queryResultsIterator.HasNext() {
		responseRange, err := queryResultsIterator.Next()
		if err != nil {
			return nil, err
		}

		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil, err
		}

		if len(compositeKeyParts) > 1 {
			returnedKey := compositeKeyParts[1]
			person, err := s.GetByNationalId(ctx, returnedKey)
			if err != nil {
				return nil, err
			}

			people = append(people, *person)
		}
	}

	if people == nil {
		return nil, fmt.Errorf("GetByFirstName: the person %s does not exist", firstName)
	}
	logger.Infoln("End: GetByFirstName called with value of: ", firstName)

	return people, nil
}

// GET BY LAST NAME AND ISMARRIED
func (s *SmartContract) GetByLastName(ctx contractapi.TransactionContextInterface,
	lastName string, isMarried bool) ([]Person, error) {

	logger := logging.NewLogger("samplechaincode")
	logger.Infoln("Start: Calling GetByFirstName function.")

	if lastName == emptyString || strconv.FormatBool(isMarried) == emptyString {
		return nil, fmt.Errorf("GetByFirstName: input parameter must not be empty")
	}

	queryResultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(lastNameAndIsMarriedIndex,
		[]string{strings.ToLower(lastName), strconv.FormatBool(isMarried)})
	if err != nil {
		return nil, err
	}
	defer queryResultsIterator.Close()

	var people []Person
	for queryResultsIterator.HasNext() {
		responseRange, err := queryResultsIterator.Next()
		if err != nil {
			return nil, err
		}

		_, compositeKeyParts, err := ctx.GetStub().SplitCompositeKey(responseRange.Key)
		if err != nil {
			return nil, err
		}

		if len(compositeKeyParts) > 1 {
			returnedKey := compositeKeyParts[2]
			person, err := s.GetByNationalId(ctx, returnedKey)
			if err != nil {
				return nil, err
			}

			people = append(people, *person)
		}
	}

	if people == nil {
		return nil, fmt.Errorf("GetByFirstName: the person %s does not exist", lastName)
	}
	logger.Infoln("End: GetByFirstName called with value of: ", lastName)

	return people, nil
}

// CHECK IF RECORD IS EXISTING IN WORLD STATE
func (s *SmartContract) IsExists(ctx contractapi.TransactionContextInterface,
	nationalId string) (bool, error) {

	logger := logging.NewLogger("samplechaincode")
	logger.Infoln("Start: Calling IsExists function.")

	queryResult, err := ctx.GetStub().GetState(nationalId)
	if err != nil {
		return false, nil
	}

	logger.Infoln("End: IsExists called with values of: ", nationalId)
	return queryResult != nil, nil
}

// GET ALL RECORDS ON THE LEDGER
func (s *SmartContract) GetAll(ctx contractapi.TransactionContextInterface) ([]QueryResult, error) {

	logger := logging.NewLogger("samplechaincode")
	logger.Infoln("Start: Calling GetAll function.")
	resultsIterator, err := ctx.GetStub().GetStateByRange(emptyString, emptyString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var people []QueryResult
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var person Person
		err = json.Unmarshal(queryResponse.Value, &person)
		if err != nil {
			return nil, err
		}

		queryResult := QueryResult{Key: queryResponse.Key, Record: &person}
		people = append(people, queryResult)
	}

	if people == nil {
		return nil, fmt.Errorf("no records available")
	}

	logger.Infoln("End: GetAll called.")
	return people, nil
}

// DELETE PERSON
func (s *SmartContract) DeletePerson(ctx contractapi.TransactionContextInterface,
	key string) error {

	logger := logging.NewLogger("transactionchaincode")
	logger.Infoln("Start: Calling DeleteFinanceRequest function.")
	data, err := s.GetByNationalId(ctx, key)
	if err != nil {
		return err
	}

	err = ctx.GetStub().DelState(key)
	logger.Infoln("DeleteFinanceRequest called with value of: ", key)
	if err != nil {
		return fmt.Errorf("failed to delete finance request %s: %v", key, err)
	}

	firstNameIndexKey, _ := ctx.GetStub().CreateCompositeKey(firstNameIndex, []string{strings.ToLower(data.FirstName), key})
	lastNameAndIsMarriedIndexKey, _ := ctx.GetStub().CreateCompositeKey(strings.ToLower(lastNameAndIsMarriedIndex),
		[]string{data.LastName, strconv.FormatBool(data.IsMarried), key})

	err = ctx.GetStub().DelState(firstNameIndexKey)
	logger.Infoln("firstNameIndexKey deleted: ", firstNameIndexKey)
	if err != nil {
		return err
	}
	logger.Infoln("lastNameAndIsMarriedIndexKey deleted: ", lastNameAndIsMarriedIndexKey)
	return ctx.GetStub().DelState(lastNameAndIsMarriedIndexKey)
}

/*
LEVEL DB IMPLEMENTATION

+++This command will initialize your project module+++
go mod init <containing folder>
go mod tidy

+++This command will bring up the network and deploy chaincode+++

./network.sh up createChannel -ca

./network.sh deployCC -ccn basic -ccp ../template/ -ccl go

export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

+++This command will invoke the chaincode+++

peer chaincode invoke \
-o localhost:7050 \
--ordererTLSHostnameOverride orderer.example.com \
--tls \
--cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem \
-C mychannel \
-n basic \
--peerAddresses localhost:7051 \
--tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt \
--peerAddresses localhost:9051 \
--tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt \
-c '{"function":"CreatePerson","Args":["{\"firstName\":\"Darwin\",\"lastName\":\"Bonsol\",\"age\":27,\"nationalId\":\"1000000000\",\"socialMediaAccounts\":[{\"websiteUrl\":\"www.test01.com\",\"dateCreated\":\"04/21/2022\"},{\"websiteUrl\":\"www.bing.com\",\"dateCreated\":\"04/27/2022\"}],\"isMarried\":false}"]}'


-c '{"function":"UpdatePerson","Args":["1000000000","Wilawan"]}'

+++This command will query the chaincode+++

peer chaincode query -C mychannel -n basic -c '{"Args":["GetByNationalId","1000000000"]}'

-c '{"Args":["GetByFirstName","Darwin"]}'

-c '{"Args":["GetByLastName","Bonsol","false"]}'
*/
