package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
    "encoding/json"
    "github.com/golang-collections/go-datastructures/queue"
    "github.com/hyperledger/fabric-chaincode-go/shim"
    pb "github.com/hyperledger/fabric-protos-go/peer"
)

type SimpleChaincode struct{
}

// Define transaction structures
type AdvertisementTransaction struct{
    TxId string                 `json:"TxId"`
    TxType string               `json:"TxType"`
	Title string 				`json:"Title"`
    Description string          `json:"Description"`
	Price string	    		`json:"Price"`
	DataType string				`json:"DataType"`
    IPAddress string			`json:"IPAddress"`
    OrgID string                `json:"OrgID"`
    TxIndex string              `json:"TxIndex"`
	//publicKey byte[]			`json:"pk"`
}

type BuyTransaction struct{
    TxId string                 `json:"TxId"`
	AdvertisementTxID string	`json:"AdvertisementTxID"`
    TxType string               `json:"TxType"`
	IPAddress string			`json:"IPAddress"`
    ClientID string             `json:"Org"`
    TxIndex string              `json:"TxIndex"`
	//publicKey byte[]			`json:"pk"`
}

// Define client (org) structure
type Client struct{
	//publicKey byte[]			`json:"pk"`i
	Assets string  				`json:"Assets"`
	OrgID string				`json:"OrgID"`
	//publicKey byte[]			`json:"pk"`i
}


// Initialize queue of pending transactions
var q queue.Queue

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple Chaincode: %s", err)
	}
}

// Initialize the smart contract with 2 organizations and their respective assets
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	var assetsOrgA, assetsOrgB string
	var err error
	var OrgA, OrgB string

	_, args := stub.GetFunctionAndParameters()

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expected 4 arguments")
	}
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	
	OrgA = args[0]
	assetsOrgA = args[1]
	OrgB = args[2]
	assetsOrgB = args[3]

    // Create OrgA JSON string to be stored in the global state
	contractJSONasString := `{"org": "` + OrgA +`","assets": "` + assetsOrgA +`"}`
	contractJSONasBytes:= []byte(contractJSONasString)

    // Save transaction to global state
	err = stub.PutState(OrgA, contractJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

    // Create OrgB JSON string to be stored in the global state
	contractJSONasString = `{"org": "` + OrgB +`","assets": "` + assetsOrgB +`"}`
	contractJSONasBytes = []byte(contractJSONasString)

    // Save transaction to global state
	err = stub.PutState(OrgB, contractJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	return shim.Success(nil)
}

// Define invocable functions on the smart contract
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "issueAdvertisement" {
		return t.issueAdvertisement(stub, args)
	} else if function == "issueBuy" {
		return t.issueBuy(stub, args)
	} else if function == "getHistoryForTransaction" {
		return t.getHistoryForTransaction(stub, args)
	} else if function == "queryAdvertisementByDataType" {
		return t.queryAdvertisementByDataType(stub,args)
	} else if function == "getAccountBalance" {
		return t.getAccountBalance(stub,args)
	} else if function == "getPendingTransactions" {
		return t.getPendingTransactions(stub,args)
	}

	return shim.Error("Received unknown function invocation")
}

// Issue a new advertisement transaction on the blockchain
func (t *SimpleChaincode) issueAdvertisement (stub shim.ChaincodeStubInterface, args []string) pb.Response{

	var err error

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expected 6 arguments. Usage: '{\"Args\":[\"<Title>\",\"<Description>\",\"<Price>\",\"<DataType>\",\"<IPAddress>\",\"<OrgID>\"]}'")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}

    txID := stub.GetTxID()
    txType := "advertisement"
	title := strings.ToLower(args[0])
    description := strings.ToLower(args[1])
	price := args[2]
	dataType := strings.ToLower(args[3])
	ipAddress:= args[4]
	orgID := args[5]

	// Check if transaction already exists in the blockchain
	contractAsBytes, err := stub.GetState(txID)
	if err != nil {
		return shim.Error("Failed to get contract: " + err.Error())
	} else if contractAsBytes != nil {
		fmt.Println("This txID already exists: " + txID)
		return shim.Error("This contract already exists: " + txID)
	}

    // Create JSON to be stored in the global state
    contractJSONasString := `{"TxID": "` + txID + `","txType": "` + txType + `","title": "` + title + `","description": "` + description + `","price": "` + price +`","dataType": "` + dataType + `","ipAddress": "` + ipAddress + `","orgID": "` + orgID + `"}`
	contractJSONasBytes:= []byte(contractJSONasString)

	// Save transaction to global state
	err = stub.PutState(txID, contractJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(txID))
}

// Issue a new buy transaction on the blockchain
func (t *SimpleChaincode) issueBuy (stub shim.ChaincodeStubInterface, args []string) pb.Response{

    var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expected 3 arguments. Usage: '{\"Args\":[\"<advertisementTransactionID>\",\"<ipAddress>\",\"<OrgID>\"]}'")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}


    txID := stub.GetTxID()
    advertisementTxID:=strings.ToLower(args[0])
    txType := "buy"
	ipAddress:= args[1]
	dstOrgID := args[2]

	// Check if transaction already exists in the blockchain
	contractAsBytes, err := stub.GetState(txID)
	if err != nil {
		return shim.Error("Failed to get contract: " + err.Error())
	} else if contractAsBytes != nil {
		fmt.Println("This txID already exists: " + txID)
		return shim.Error("This contract already exists: " + txID)
	}

    // Get the corresponding advertisement transaction from the global state
	adContractAsBytes, err := stub.GetState(advertisementTxID)
	if err != nil {
		return shim.Error("Failed to get contract: " + err.Error())
	} else if adContractAsBytes == nil {
		fmt.Println("The referenced advertisement transaction " + advertisementTxID + " does not exist.")
		return shim.Error("The referenced advertisement transaction " + advertisementTxID + " does not exist.")
	}
	referencedAdvertisement := AdvertisementTransaction{}
	err = json.Unmarshal(adContractAsBytes, &referencedAdvertisement)
	if err != nil {
		return shim.Error(err.Error())
	}

    // Extract relevant info from the referenced transaction
	advertisementTxId := referencedAdvertisement.TxId
    srcOrgID := referencedAdvertisement.OrgID
    dataPrice, _ := strconv.Atoi(referencedAdvertisement.Price)

    // Retrieve the involved orgs and their respective assets
	srcOrg := Client{}
	SrcOrgAssetsbytes, err := stub.GetState(srcOrgID)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if SrcOrgAssetsbytes == nil {
		return shim.Error("Entity not found")
	}
	err = json.Unmarshal(SrcOrgAssetsbytes, &srcOrg)
	if err != nil {
		return shim.Error(err.Error())
	}
    srcOrgAssets, _ := strconv.Atoi(srcOrg.Assets)

	dstOrg := Client{}
	DstOrgAssetsbytes, err := stub.GetState(dstOrgID)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if DstOrgAssetsbytes == nil {
		return shim.Error("Entity not found")
	}
	err = json.Unmarshal(DstOrgAssetsbytes, &dstOrg)
	if err != nil {
		return shim.Error(err.Error())
	}
    dstOrgAssets, _ := strconv.Atoi(dstOrg.Assets)

    // Transfer assets from buyer to seller
	srcOrgAssets = srcOrgAssets + dataPrice
	dstOrgAssets = dstOrgAssets - dataPrice

    // Abort transaction if buyer does not have enough funds
	if dstOrgAssets < 0.0 {
		return shim.Error("The source organization does not have enough assets to conclude this transaction!")
	}

	assetsSrcOrgAsString := strconv.Itoa(srcOrgAssets)
	assetsDstOrgAsString := strconv.Itoa(dstOrgAssets)

    // Create JSON string and update orgs assets on the global state
	clientAJSONasString := `{"org": "` + srcOrgID +`","assets": "` + assetsSrcOrgAsString +`"}`
	clientAJSONasBytes:= []byte(clientAJSONasString)

	err = stub.PutState(srcOrgID, clientAJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	clientBJSONasString := `{"org": "` + dstOrgID +`","assets": "` + assetsDstOrgAsString +`"}`
	clientBJSONasBytes:= []byte(clientBJSONasString)

	err = stub.PutState(dstOrgID, clientBJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

    // Create JSON string of the buy transaction and store it in the global state
    contractJSONasString := `{"TxID": "` + txID + `", "advertisementTxId": "` + advertisementTxId + `","txType": "` + txType + `","ipAddress": "` + ipAddress + `","orgID": "` + srcOrgID + `"}`
	contractJSONasBytes:= []byte(contractJSONasString)

	err = stub.PutState(txID, contractJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

    // Add buy transaction to queue for controller processing
	queueJSONasString := `{"TxId": "` + txID +`","SrcIPAddress": "` + ipAddress +`","DstIPAddress": "` + referencedAdvertisement.IPAddress +`"}`
    err = q.Put(queueJSONasString)

	return shim.Success([]byte(txID))

}

// Returns all unprocessed buy transactions
func (t *SimpleChaincode) getPendingTransactions (stub shim.ChaincodeStubInterface, args []string) pb.Response {

    if len(args) < 0 {
        return shim.Error("Incorrect number of arguments. Expected 0 arguments")
    }

    // If queue is empty, return error
    if (q.Len() == 0){
        return shim.Error("No buy transactions to be processed!")
    }

    // Get transactions from the queue and return them in JSON format
    results, err := q.Get(q.Len())
    if err != nil {
        return shim.Error(err.Error())
    }

    resultsAsString := strings.ReplaceAll(fmt.Sprintf("%v",results),"} {","}, {")

    return shim.Success([]byte(resultsAsString))
}


// Return the account balance for a user in the system
func (t *SimpleChaincode) getAccountBalance (stub shim.ChaincodeStubInterface, args []string) pb.Response{
    
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	OrgID := args[0]
    result, err := stub.GetState(OrgID)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(result)
}

// Track a transaction history by its txID
func (t *SimpleChaincode) getHistoryForTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	txID := args[0]

    // Retrieve transaction history
	resultsIterator, err := stub.GetHistoryForKey(txID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// Format results as a JSON array containing historic values for the the transaction
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

func (t *SimpleChaincode) queryAdvertisementByDataType (stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	dataType := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"dataType\":\"%s\"}}", dataType)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// Increase funds for an organization
func (t *SimpleChaincode) addAssetsToOrganization(stub shim.ChaincodeStubInterface, args[] string) pb.Response{

	var OrgA string // Entities
	var amount int
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting txID of the organizations and amount to be added")
	}

	OrgA = args[0]
	amount, err = strconv.Atoi(args[1])
	if err != nil{
		return shim.Error("Error converting amount to integer: " + err.Error())
	}

	// Get the state from the ledger
	organization := Client{}
	OrgAbytes, err := stub.GetState(OrgA)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + OrgA + "\"}"
		return shim.Error(jsonResp)
	}
	err = json.Unmarshal(OrgAbytes, &organization)
	if err != nil {
		return shim.Error(err.Error())
	}
	if OrgAbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + OrgA + "\"}"
		return shim.Error(jsonResp)
	}

    assetsAsInt, _ := strconv.Atoi(organization.Assets)
    assetsAsInt = assetsAsInt + amount
	assetsOrgA := strconv.Itoa(assetsAsInt) 

	clientJSONasString := `{"org": "` + OrgA +`","assets": "` + assetsOrgA +`"}`
	clientJSONasBytes:= []byte(clientJSONasString)

	err = stub.PutState(OrgA, clientJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
