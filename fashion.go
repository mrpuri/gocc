package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)
//represents the chaincode object referenced throughout
type fashionChaincode struct{

}
func (fashion *fashionChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("init executed")

	return shim.Success(nil)
}

func (fashion *fashionChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("invoke executed")
	payload := []byte("This is the payload. ")
	//return shim.Success(payload) utility function for generating the response
	//instance of the response structure which is similar to http response                                                               
	return peer.Response{Status:401, Message: "Unauthorized", Payload: payload}
}

func main() {
	fmt.Println("Started Chaincode")

	err := shim.Start(new(fashionChaincode))
	if err !=  nil {
		fmt.Println("Error starting chaincode : %s", err)
	}
}