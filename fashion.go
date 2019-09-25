package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

//ChaincodeName to create an instance of logger
const ChaincodeName = "fashion"

var logger = shim.NewLogger(ChaincodeName)

//represents the chaincode object referenced throughout
//reciever for chaincode shim functions
type fashionChaincode struct {
}

type cloth struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Type       string `json:"type"`    //the fieldtags are needed to keep case from bouncing around
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
	// logger.Debug("Invoke executed for log")
	// //Payload := []byte("This is the payload. ")
	// fmt.Println("transaction id", stub.GetTxID())
	// t, _ := stub.GetTxTimestamp()
	// TxTimestamp := time.Unix(t.GetSeconds(), 0)
	// fmt.Println("transaction time stamp", TxTimestamp)
	// fmt.Println("the channel id is", stub.GetChannelID())

	argsArray := stub.GetArgs()
	fmt.Println("get args byte array output")
	for ndx, arg := range argsArray {
		argStr := string(arg)
		fmt.Printf("[%d] = %s", ndx, argStr)
	}
	
	fmt.Println(" the output of the get args function")
	fmt.Println(stub.GetStringArgs())

	fmt.Println("output of functions and parameters")
	funcName, args := stub.GetFunctionAndParameters()
	fmt.Printf("function name = %s \n Args = %s \n", funcName, args)

	fmt.Println("getting arguments of slice")
	argsSlice, _ := stub.GetArgsSlice()
	length := len(argsSlice)
	fmt.Println(length, argsSlice)
	return shim.Success(nil) //utility function for generating the response

	//instance of the response structure which is similar to http response
	//return peer.Response{Status:401, Message: "Unauthorized", Payload: payload}
	//sending success and error response.
}

func main() {
	fmt.Println("Started Chaincode")

	err := shim.Start(new(fashionChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode : %v", err)
	}
}
