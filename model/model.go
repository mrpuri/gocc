package model

//FashionChaincode structure -- instance of the chaincode
//represents the chaincode object referenced throughout, reciever for chaincode shim functions
type GFashionChaincode struct {
	
}

//Cloth structure -- example for asset properties
type Cloth struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"type"`    //the fieldtags are needed to keep case from bouncing around
	Color      string `json:"color"`
	Size       int    `json:"size"`
	Owner      string `json:"owner"`
}
