package controller

import (
	model "fashion/model"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type cFashionChaincode struct {
	model.GFashionChaincode
}

//To be included in the invoke/ init/ set function
// //testing get and set functions for state data
// valuet, err := stub.GetState("token")
// if err != nil {
// 	return shim.Error(err.Error())
// }
// intval, err := strconv.Atoi(string(valuet))
// if err != nil {
// 	return shim.Success([]byte("false"))
// }
// intval += 10

// stub.PutState("token", []byte(strconv.Itoa(intval)))

// //testing ends here

//for testing token added to the chaincode state, retrieved through the get function -- CLI ->'{"Args":["get"]}'

//Get function exported to fashion.go
func Get(stub shim.ChaincodeStubInterface) peer.Response {
	var token string
	var val []byte
	var err error

	if val, err = stub.GetState("token"); err != nil {
		fmt.Println("failed to get the token")
		return shim.Error("get failed!")
	}
	if val == nil {
		token = "-1"
	} else {
		token = "token =" + string(val)
	}

	return shim.Success([]byte(token))

}
