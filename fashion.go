package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//ChaincodeName to create an instance of logger
const ChaincodeName = "fashion"

var logger = shim.NewLogger(ChaincodeName)

//represents the chaincode object referenced throughout, reciever for chaincode shim functions
type fashionChaincode struct {
}

type cloth struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"type"`    //the fieldtags are needed to keep case from bouncing around
	Color      string `json:"color"`
	Size       int    `json:"size"`
	Owner      string `json:"owner"`
}

func (fashion *fashionChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("init executed")
	logger.Debug("Init executed for log")

	return shim.Success(nil)
}

func (fashion *fashionChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("invoke executed")
	// {logger.Debug("Invoke executed for log")
	// //Payload := []byte("This is the payload. ")
	// fmt.Println("transaction id", stub.GetTxID())
	// t, _ := stub.GetTxTimestamp()
	// TxTimestamp := time.Unix(t.GetSeconds(), 0)
	// fmt.Println("transaction time stamp", TxTimestamp)
	// fmt.Println("the channel id is", stub.GetChannelID())

	// argsArray := stub.GetArgs()
	// fmt.Println("get args byte array output")
	// for ndx, arg := range argsArray {
	// 	argStr := string(arg)
	// 	fmt.Printf("[%d] = %s", ndx, argStr)
	// }

	// fmt.Println(" the output of the get args function")
	// fmt.Println(stub.GetStringArgs())

	// fmt.Println("output of functions and parameters")
	// funcName, args := stub.GetFunctionAndParameters()
	// fmt.Printf("function name = %s \n Args = %s \n", funcName, args)

	// fmt.Println("getting arguments of slice")
	// argsSlice, _ := stub.GetArgsSlice()
	// length := len(argsSlice)
	// fmt.Println(length, argsSlice)
	// return shim.Success(nil) //utility function for generating the response

	//instance of the response structure which is similar to http response
	//return peer.Response{Status:401, Message: "Unauthorized", Payload: payload}
	//sending success and error response.
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initcloth" { //create a new clothing asset
		return fashion.initcloth(stub, args)
	} //else if function == "transferMarble" { //change owner of a specific marble
	// 	return fashion.transferMarble(stub, args)
	// } else if function == "transferMarblesBasedOnColor" { //transfer all marbles of a certain color
	// 	return t.transferMarblesBasedOnColor(stub, args)
	// } else if function == "delete" { //delete a marble
	// 	return t.delete(stub, args)
	// } else if function == "readMarble" { //read a marble
	// 	return t.readMarble(stub, args)
	// } else if function == "queryMarblesByOwner" { //find marbles for owner X using rich query
	// 	return t.queryMarblesByOwner(stub, args)
	// } else if function == "queryMarbles" { //find marbles based on an ad hoc rich query
	// 	return t.queryMarbles(stub, args)
	// } else if function == "getHistoryForMarble" { //get history of values for a marble
	// 	return t.getHistoryForMarble(stub, args)
	// } else if function == "getMarblesByRange" { //get marbles based on range query
	// 	return t.getMarblesByRange(stub, args)
	// }

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")

}

// ============================================================
// initcloth - create a new clothing asset, stores into chaincode state
// ============================================================
func (fashion *fashionChaincode) initcloth(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error

	//   0       1       2     3
	// "shirt", "blue", "40", "bob"
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init fashion")
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
	clothName := args[0]
	color := strings.ToLower(args[1])
	owner := strings.ToLower(args[3])
	size, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}

	// ==== Check if marble already exists ====
	clothAsBytes, err := stub.GetState(clothName)
	if err != nil {
		return shim.Error("Failed to get cloths: " + err.Error())
	} else if clothAsBytes != nil {
		fmt.Println("This cloth already exists: " + clothName)
		return shim.Error("This clothing already exists: " + clothName)
	}

	// ==== Create cloth object and marshal to JSON ====
	objectType := "Clothing"
	cloth := &cloth{objectType, clothName, color, size, owner}
	clothJSONasBytes, err := json.Marshal(cloth)
	if err != nil {
		return shim.Error(err.Error())
	}
	//Alternatively, build the cloth json string manually if you don't want to use struct marshalling
	//marbleJSONasString := `{"docType":"Marble",  "name": "` + marbleName + `", "color": "` + color + `", "size": ` + strconv.Itoa(size) + `, "owner": "` + owner + `"}`
	//marbleJSONasBytes := []byte(str)

	// === Save marble to state ===
	err = stub.PutState(clothName, clothJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//  ==== Index the marble to enable color-based range queries, e.g. return all blue marbles ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~color~name.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexName := "color~name"
	colorNameIndexKey, err := stub.CreateCompositeKey(indexName, []string{cloth.Color, cloth.Name})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(colorNameIndexKey, value)

	// ==== Marble saved and indexed. Return success ====
	fmt.Println("- end init cloth the data added is: ")
	fmt.Println(cloth)
	return shim.Success(nil)
}

func main() {
	fmt.Println("Started Chaincode")

	err := shim.Start(new(fashionChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode : %v", err)
	}
}
