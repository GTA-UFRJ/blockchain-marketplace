package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
    "encoding/json"

   "github.com/hyperledger/fabric-chaincode-go/shim"
   pb "github.com/hyperledger/fabric-protos-go/peer"
)

type SimpleChaincode struct{
}

// Define transaction structures
// TODO: Verify if there are better types than "string" for the transaction fields
type AdvertisementTransaction struct{
    TxId string                 `json:"TxId"`
    TxType string               `json:"TxType"`
	Name string 				`json:"Name"`
	Price string				`json:"Price"`
	DataType string				`json:"DataType"`
    IPAddress string			`json:"IPAddress"`
    ClientID string             `json:"Org"`
    TxIndex int                 `json:"TxIndex"`
	//publicKey byte[]			`json:"pk"`
}

type BuyTransaction struct{
    TxId string                 `json:"TxId"`
	AdvertisementTxID string	`json:"AdvertisementTxID"`
	//Price string				`json:"Price"`
    TxType string               `json:"TxType"`
	IPAddress string			`json:"IPAddress"`
    ClientID string             `json:"Org"`
    TxIndex int                 `json:"TxIndex"`
	//publicKey byte[]			`json:"pk"`
}

type Client struct{
	//publicKey byte[]			`json:"pk"`i
    ClientID string             `json:"ClientID"`
	Assets int  				`json:"Assets"`
	Org string					`json:"Org"`
}


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

    lastTxIndex := 0

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


	/*assetsOrgA, err = strconv.Atoi(args[1])
	if err != nil{
		return shim.Error(err.Error())
	}
	
	assetsOrgB = strconv.Atoi(args[3])
	if err != nil{
		return shim.Error(err.Error())
	}*/

	fmt.Printf("assetsOrgA:%d, assetsOrgB:%d", assetsOrgA, assetsOrgB)
	contractJSONasString := `{"org": "` + OrgA +`","assets": "` + assetsOrgA +`"}`
	contractJSONasBytes:= []byte(contractJSONasString)


	// === Save transaction to state ===
	err = stub.PutState(OrgA, contractJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	contractJSONasString = `{"org": "` + OrgB +`","assets": "` + assetsOrgB +`"}`
	contractJSONasBytes = []byte(contractJSONasString)


	// === Save transaction to state ===
	err = stub.PutState(OrgB, contractJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

    // Save lastTxIndex to state
	err = stub.PutState("lastTxIndex", lastTxIndex)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	return shim.Success(nil)
}

// Define invocable functions on the smart contract
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()

	if function == "issueAdvertisement" { //create an advertisement transaction
		return t.issueAdvertisement(stub, args)
	} else if function == "issueBuy" {
		return t.issueBuy(stub, args)
	} else if function == "getHistoryForTransaction" { //get history of values for a transaction
		return t.getHistoryForTransaction(stub, args)
	} else if function == "queryAdvertisementByDataType"{
		return t.queryAdvertisementByDataType(stub,args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")

}

// Issue a new advertisement transaction on the blockchain
func (t *SimpleChaincode) issueAdvertisement (stub shim.ChaincodeStubInterface, args []string) pb.Response{

	var err error

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Exepected 5 arguments. Usage: '{\"Args\":[\"<Name>\",\"<Price>\",\"<DataType>\",\"<IPAddress>\",\"<ClientID>\"]}'")
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

	//transactionIssuer, err := stub.GetCreator()

    txID := stub.GetTxID()
    txType := "advertisement"
	name := strings.ToLower(args[0])
	price := args[1]
	dataType := strings.ToLower(args[2])
	ipAddress:= args[3]
	clientID := args[4]
    txIndex :=  stub.GetState("lastTxIndex")+1

	// ==== Check if transaction already exists ====
	contractAsBytes, err := stub.GetState(name)
	if err != nil {
		return shim.Error("Failed to get contract: " + err.Error())
	} else if contractAsBytes != nil {
		fmt.Println("This name already exists: " + name)
		return shim.Error("This contract already exists: " + name)
	}

    contractJSONasString := `{"TxID": "` + txID + `","txType": "` + txType + `","name": "` + name + `","price": "` + price +`","dataType": "` + dataType + `","ipAddress": "` + ipAddress + `","clientID": "` + clientID + `","txIndex": "` + txIndex + `"}`
    //contractJSONasString := `{"name": "` + name +`","price": "` + price +`","dataType": "` + dataType +`","TxId": "` + transactionTxId + `"}`
	contractJSONasBytes:= []byte(contractJSONasString)


	// === Save transaction to state ===
	err = stub.PutState(name, contractJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}


	// ==== Transaction saved. Return success ====
	return shim.Success(nil)
}

// Issue a new buy transaction on the blockchain
func (t *SimpleChaincode) issueBuy (stub shim.ChaincodeStubInterface, args []string) pb.Response{
	var err error
	var OrgA, OrgB string    // Entities
	var assetsOrgA, assetsOrgB int // Asset holdings
	//var price int          // Transaction value



	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Exepected 4 arguments. Usage: '{\"Args\":[\"<transactionName>\",\"<advertisementTransactionName>\",\"<ipAdress>\", \"<price>\",\"<srcOrganization>\",\"<destOrganization>\"]}'")
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



	//transactionIssuer, err := stub.GetCreator()

	name := strings.ToLower(args[0])
	adveritisementTransactionName := strings.ToLower(args[1])
	ipAddress := args[2]
	price := args[3]

	// ==== Check if transaction already exists ====
	contractAsBytes, err := stub.GetState(name)
	if err != nil {
		return shim.Error("Failed to get contract: " + err.Error())
	} else if contractAsBytes != nil {
		fmt.Println("This name already exists: " + name)
		return shim.Error("This contract already exists: " + name)
	}

	adContractAsBytes, err := stub.GetState(adveritisementTransactionName)
	if err != nil {
		return shim.Error("Failed to get contract: " + err.Error())
	} else if adContractAsBytes == nil {
		fmt.Println("The referenced advertisement transaction" + adveritisementTransactionName + "does not exist")
		return shim.Error("The referenced advertisement transaction" + adveritisementTransactionName + "does not exist")
	}


	referencedAdvertisement := AdvertisementTransaction{}
	err = json.Unmarshal(adContractAsBytes, &referencedAdvertisement)
	if err != nil {
		return shim.Error(err.Error())
	}

	advertisementTxId := referencedAdvertisement.TxId

	buyerOffer, err := strconv.Atoi(price)
	if err != nil {
		return shim.Error("Could not convert the price informed to float: " + err.Error())
	}


	advertisementPrice, err := strconv.Atoi(referencedAdvertisement.Price)
	if err != nil {
		return shim.Error("Could not convert the referenced advertisement price to float: " + err.Error())
	}

	if buyerOffer < advertisementPrice {
		return shim.Error("Price payed must be equal or higher to the advertised price!\n")
	}

	contractJSONasString := `{"name": "` + name +`","price": "` + price +`","ipAddress": "` + ipAddress +`","advertisementTxId": "` + advertisementTxId +`"}`
	contractJSONasBytes:= []byte(contractJSONasString)


	
	OrgA = args[4]
	OrgB = args[5]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	organization := Client{}
	OrgAAssetsbytes, err := stub.GetState(OrgA)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if OrgAAssetsbytes == nil {
		return shim.Error("Entity not found")
	}
	err = json.Unmarshal(OrgAAssetsbytes, &organization)
	if err != nil {
		return shim.Error(err.Error())
	}
	assetsOrgA, _ = strconv.Atoi(organization.Assets)

	organizationB := Client{}
	OrgBAssetsbytes, err := stub.GetState(OrgB)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if OrgBAssetsbytes == nil {
		return shim.Error("Entity not found")
	}
	err = json.Unmarshal(OrgBAssetsbytes, &organizationB)
	if err != nil {
		return shim.Error(err.Error())
	}
	assetsOrgB, err = strconv.Atoi(organizationB.Assets)
	if err != nil {
		return shim.Error("Failed to convert string to int: " + err.Error())
	}

	assetsOrgA = assetsOrgA - buyerOffer
	assetsOrgB = assetsOrgB + buyerOffer
	
	if assetsOrgA < 0.0{
		shim.Error("ClientA does not have money to conclude this transaction!\n")
	}

	fmt.Printf("assetsOrgA = %d, assetsOrgB = %d\n", assetsOrgA, assetsOrgB)

	assetsOrgAAsString := strconv.Itoa(assetsOrgA)
	assetsOrgBAsString := strconv.Itoa(assetsOrgB)

	clientAJSONasString := `{"org": "` + OrgA +`","assets": "` + assetsOrgAAsString +`"}`
	clientAJSONasBytes:= []byte(clientAJSONasString)


	// === Save transaction to state ===
	err = stub.PutState(OrgA, clientAJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	clientBJSONasString := `{"org": "` + OrgB +`","assets": "` + assetsOrgBAsString +`"}`
	clientBJSONasBytes:= []byte(clientBJSONasString)


	// === Save transaction to state ===
	err = stub.PutState(OrgB, clientBJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}



	// === Save transaction to state ===
	err = stub.PutState(name, contractJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}


	// ==== Transaction saved. Return success ====
	return shim.Success(nil)

}

// Track a transaction history by its name
// TODO: replace name with TxID
func (t *SimpleChaincode) getHistoryForTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	name := args[0]

	resultsIterator, err := stub.GetHistoryForKey(name)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the the transaction
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
		return shim.Error("Incorrect number of arguments. Expecting name of the organizations and amount to be added")
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

	OrgAAssetsAsInt, err := strconv.Atoi(organization.Assets)
	if err != nil{
		return shim.Error("Error converting amount to integer: " + err.Error())
	}

	assetsAsInt := OrgAAssetsAsInt + amount
	
	assetsOrgA := strconv.Itoa(assetsAsInt) 

	clientJSONasString := `{"org": "` + OrgA +`","assets": "` + assetsOrgA +`"}`
	clientJSONasBytes:= []byte(clientJSONasString)


	// === Save transaction to state ===
	err = stub.PutState(OrgA, clientJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

		

	return shim.Success(nil)

}

