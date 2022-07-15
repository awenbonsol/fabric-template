package main_test

import (
	"encoding/json"
	"fmt"
	"templatecouchdb"
	"templatecouchdb/mocks"
	"testing"

	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	"github.com/stretchr/testify/require"
)

func TestCreatePersonPositiveScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	samplechaincode := main.SmartContract{}
	_, err := samplechaincode.CreatePerson(transactionContext, mockPerson)
	require.NoError(t, err)
}

func TestCreatePersonNegativeScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	samplechaincode := main.SmartContract{}
	_, err := samplechaincode.CreatePerson(transactionContext, mockPerson)
	require.EqualError(t, err, "national id 100000000000000 must not exceed 12 characters")
}

func TestUpdatePersonPositiveScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	bytes, err := json.Marshal(mockPerson)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	samplechaincode := main.SmartContract{}
	_, err = samplechaincode.UpdatePerson(transactionContext, "100000000000", "Driodoco")
	require.NoError(t, err)
}

func TestUpdatePersonNegativeScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	_, err := json.Marshal(mockPerson)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("failed retrieving all finance request"))
	samplechaincode := main.SmartContract{}
	_, err = samplechaincode.UpdatePerson(transactionContext, "100000000000", "Driodoco")
	require.EqualError(t, err, "GetByNationalId: failed to read from world state: failed retrieving all finance request")

	chaincodeStub.GetStateReturns(nil, nil)
	_, err = samplechaincode.UpdatePerson(transactionContext, "100000000000", "Driodoco")
	require.EqualError(t, err, "GetByNationalId: the person 100000000000 does not exist")
}

func TestGetByNationalIdPositiveScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	bytes, err := json.Marshal(mockPerson)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(bytes, nil)
	samplechaincode := main.SmartContract{}
	result, err := samplechaincode.GetByNationalId(transactionContext, "100000000000")
	require.NoError(t, err)
	require.Equal(t, *result, mockPerson)
}

func TestGetByNationalIdNegativeScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	_, err := json.Marshal(mockPerson)
	require.NoError(t, err)

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("failed retrieving all finance request"))
	samplechaincode := main.SmartContract{}
	_, err = samplechaincode.GetByNationalId(transactionContext, "100000000000")
	require.EqualError(t, err, "GetByNationalId: failed to read from world state: failed retrieving all finance request")

	chaincodeStub.GetStateReturns(nil, nil)
	_, err = samplechaincode.GetByNationalId(transactionContext, "100000000000")
	require.EqualError(t, err, "GetByNationalId: the person 100000000000 does not exist")
}

func TestGetByFirstNamePositiveScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	bytes, err := json.Marshal(mockPerson)
	require.NoError(t, err)

	iterator := &mocks.StateQueryIterator{}
	iterator.HasNextReturnsOnCall(0, true)
	iterator.HasNextReturnsOnCall(1, false)
	iterator.NextReturns(&queryresult.KV{Value: bytes}, nil)
	chaincodeStub.GetQueryResultReturns(iterator, nil)

	samplechaincode := main.SmartContract{}

	chaincodeStub.GetStateReturns(bytes, nil)
	_, err = samplechaincode.GetByFirstName(transactionContext, "Darwin")
	require.NoError(t, err)
}

func TestGetByFirstNameNegativeScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	bytes, err := json.Marshal(mockPerson)
	require.NoError(t, err)

	iterator := &mocks.StateQueryIterator{}
	iterator.HasNextReturnsOnCall(0, true)
	iterator.HasNextReturnsOnCall(1, false)
	iterator.NextReturns(&queryresult.KV{Value: bytes}, nil)
	chaincodeStub.GetQueryResultReturns(iterator, nil)

	samplechaincode := main.SmartContract{}

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("failed retrieving all finance request"))
	_, err = samplechaincode.GetByNationalId(transactionContext, "100000000000")
	require.EqualError(t, err, "GetByNationalId: failed to read from world state: failed retrieving all finance request")

	chaincodeStub.GetStateReturns(nil, nil)
	_, err = samplechaincode.GetByNationalId(transactionContext, "100000000000")
	require.EqualError(t, err, "GetByNationalId: the person 100000000000 does not exist")
}

func TestGetByLastNamePositiveScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	bytes, err := json.Marshal(mockPerson)
	require.NoError(t, err)

	iterator := &mocks.StateQueryIterator{}
	iterator.HasNextReturnsOnCall(0, true)
	iterator.HasNextReturnsOnCall(1, false)
	iterator.NextReturns(&queryresult.KV{Value: bytes}, nil)
	chaincodeStub.GetQueryResultReturns(iterator, nil)

	samplechaincode := main.SmartContract{}

	chaincodeStub.GetStateReturns(bytes, nil)
	_, err = samplechaincode.GetByLastName(transactionContext, "Darwin", false)
	require.NoError(t, err)
}

func TestGetByLastNameNegativeScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	bytes, err := json.Marshal(mockPerson)
	require.NoError(t, err)

	iterator := &mocks.StateQueryIterator{}
	iterator.HasNextReturnsOnCall(0, true)
	iterator.HasNextReturnsOnCall(1, false)
	iterator.NextReturns(&queryresult.KV{Value: bytes}, nil)
	chaincodeStub.GetQueryResultReturns(iterator, nil)

	samplechaincode := main.SmartContract{}

	chaincodeStub.GetStateReturns(nil, fmt.Errorf("failed retrieving all finance request"))
	_, err = samplechaincode.GetByNationalId(transactionContext, "100000000000")
	require.EqualError(t, err, "GetByNationalId: failed to read from world state: failed retrieving all finance request")

	chaincodeStub.GetStateReturns(nil, nil)
	_, err = samplechaincode.GetByNationalId(transactionContext, "100000000000")
	require.EqualError(t, err, "GetByNationalId: the person 100000000000 does not exist")
}

func TestGetAllPositiveScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	bytes, err := json.Marshal(mockPerson)
	require.NoError(t, err)

	iterator := &mocks.StateQueryIterator{}
	iterator.HasNextReturnsOnCall(0, true)
	iterator.HasNextReturnsOnCall(1, false)
	iterator.NextReturns(&queryresult.KV{Value: bytes}, nil)

	chaincodeStub.GetStateByRangeReturns(iterator, nil)
	samplechaincode := main.SmartContract{}
	result, err := samplechaincode.GetAll(transactionContext)
	require.NoError(t, err)
	require.NotNil(t, result)
}

func TestGetAllNegativeScenario(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	mockSocialMediaAccount01 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test.com",
		DateCreated: "05/05/2022",
	}

	mockSocialMediaAccount02 := main.SocialMediaAccounts{
		WebsiteUrl:  "www.test01.com",
		DateCreated: "05/10/2022",
	}

	mockSocialMediaAccounts := []main.SocialMediaAccounts{mockSocialMediaAccount01, mockSocialMediaAccount02}

	mockPerson := main.Person{
		FirstName:           "Darwin",
		LastName:            "Bonsol",
		Age:                 28,
		NationalID:          "100000000000",
		SocialMediaAccounts: mockSocialMediaAccounts,
		IsMarried:           false,
	}

	bytes, err := json.Marshal(mockPerson)
	require.NoError(t, err)

	iterator := &mocks.StateQueryIterator{}
	iterator.HasNextReturnsOnCall(0, true)
	iterator.HasNextReturnsOnCall(1, false)
	iterator.NextReturns(&queryresult.KV{Value: bytes}, nil)

	samplechaincode := main.SmartContract{}
	chaincodeStub.GetStateByRangeReturns(iterator, fmt.Errorf("no available records"))
	_, err = samplechaincode.GetAll(transactionContext)
	require.Error(t, err)
	require.EqualError(t, err, "no available records")
}
