package main
 
import ( 
 	"encoding/json" 
 	"fmt" 
	"errors"
  
 	"github.com/hyperledger/fabric/core/chaincode/shim" 
 ) 

// SimpleChaincode example simple Chaincode implementation 
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
 
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    return nil, nil
}
 
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    
	fmt.Println("Executing function " + function) 
    if function == "getShipmentDetails" {                            //read a variable 
         return t.GetShipmentDetails(stub, args)  
     } 
    fmt.Println("query did not find the func: " + function)
	return nil, nil
}
 
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
   
	
	fmt.Println("invoke did not find func: " + function) 
 
    return nil, errors.New("Received unknown function invocation") 
}
 
func (t *SimpleChaincode) GetShipmentDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	shipperBytes, err := stub.GetState(args[0])
	var shipperInformation ShipperInformation
	err = json.Unmarshal(shipperBytes, &shipperInformation)
	fmt.Println(shipperInformation.Origin)
    return shipperBytes, err
 }

 
func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Println("Could not start SimpleChaincode")
    } else {
        fmt.Println("SimpleChaincode successfully started")
    }
 
}