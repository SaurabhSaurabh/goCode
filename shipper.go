package main

import ( 
 	"encoding/json" 
 	"fmt" 
	"errors"
  
 	"github.com/hyperledger/fabric/core/chaincode/shim" 
 ) 

type SimpleChaincode struct { 
} 

type ShipperInformation struct {
    ShipmentId          	string        `json:"shipmentid"`
    Carrier             	string        `json:"carrier"`
	Shipper             	string        `json:"shipper"`
    Consignee           	string        `json:"consignee"`
    Comodity            	string        `json:"comodity"`
    ContainerRequirement	string        `json:"containerrequirement"`
    Origin             		string		  `json:"origin"`
    Destination          	string		  `json:"destination"`
}
 
func (t *ShipperInformation) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    
	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
	}
	
	shipmentData1 := args[0] // shipment id
	shipmentData2 := args[1] // carrier
	shipmentData3 := args[2] // shipper
	shipmentData4 := args[3] // consignee
	shipmentData5 := args[4] // comodity
	shipmentData6 := args[5] // container requirement
	shipmentData7 := args[6] // origin
	shipmentData8 := args[7] // destination

	ShipperInformation := ShipperInformation{
		ShipmentId:  shipmentData1,
		Carrier:    shipmentData2,
		Shipper: shipmentData3,
		Consignee: shipmentData4,
		Comodity: shipmentData5,
		ContainerRequirement: shipmentData6,
		Origin: shipmentData7,
		Destination: shipmentData8,		
	}
	
	fmt.Println(ShipperInformation)
	
	bytes, err := json.Marshal(ShipperInformation)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}	
	fmt.Println(bytes)
	
	err = stub.PutState(shipmentData1, bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}
	
	return nil, nil
}
 
func (t *ShipperInformation) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    
	fmt.Println("Executing function " + function) 
    if function == "getShipmentDetails" {                            //read a variable 
         return t.GetShipmentDetails(stub, args)  
     } 
    fmt.Println("query did not find the func: " + function)
	return nil, nil
}
 
func (t *ShipperInformation) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
   
	fmt.Println("Executing function " + function) 
    if function == "writeShipperDetails" {                            //read a variable 
         return t.WriteShipperDetails(stub, args)  
     } 
   	fmt.Println("invoke did not find func: " + function) 
 
    return nil, errors.New("Received unknown function invocation") 
}
 
func (t *ShipperInformation) WriteShipperDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
	}
	
	shipmentData1 := args[0] // shipment id
	shipmentData2 := args[1] // carrier
	shipmentData3 := args[2] // shipper
	shipmentData4 := args[3] // consignee
	shipmentData5 := args[4] // comodity
	shipmentData6 := args[5] // container requirement
	shipmentData7 := args[6] // origin
	shipmentData8 := args[7] // destination

	ShipperInformation := ShipperInformation{
		ShipmentId:  shipmentData1,
		Carrier:    shipmentData2,
		Shipper: shipmentData3,
		Consignee: shipmentData4,
		Comodity: shipmentData5,
		ContainerRequirement: shipmentData6,
		Origin: shipmentData7,
		Destination: shipmentData8,		
	}
	
	fmt.Println(ShipperInformation)
	
	bytes, err := json.Marshal(ShipperInformation)
	if err != nil {
		fmt.Println("Error marsalling")
		return nil, errors.New("Error marshalling")
	}	
	fmt.Println(bytes)
	
	err = stub.PutState(shipmentData1, bytes)
	if err != nil {
		fmt.Println("Error writing state")
		return nil, err
	}
	
	return nil, nil
}
 
func (t *ShipperInformation) GetShipmentDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 4")
	}
	shipperBytes, err := stub.GetState(args[0])
	var shipperInformation ShipperInformation
	err = json.Unmarshal(shipperBytes, &shipperInformation)
	fmt.Println(shipperInformation.Origin)
	
	shipperBytes, err = json.Marshal(shipperInformation)
	if err != nil {
		fmt.Println("Error marshaling ShipperInformation")
		return nil, errors.New("Error marshaling ShipperInformation")
	}

	fmt.Println(shipperBytes)
	
    return shipperBytes, err
 }

 
func main() {
    err := shim.Start(new(ShipperInformation))
    if err != nil {
        fmt.Println("Could not start SimpleChaincode")
    } else {
        fmt.Println("SimpleChaincode successfully started")
    }
 
}