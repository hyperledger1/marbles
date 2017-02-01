/* SUPPLY CHAIN MANAGEMENT - RETAIL SECTOR*/
package main

import (
	"errors"
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strings"
	"time"
	"strconv"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

				

type Lee  struct{
	
	Color string `json:"color"`                                     // blue, brown etc
	Size int `json:"size"`                                          // waist size 32,34 etc
}

type currency struct{
supplycoin float32
}

/*
type Description struct{
	Color string `json:"color"`
	Size int `json:"size"`
	Brand string `json:"brand"`
}

*/

type assets struct{
	 User string `json:"user"`  
         Quantity int  `json:"quantity"`
	 Typeofasset  string `json:"typeofasset"`
}


/*
type Order struct{
	OrderID string `json:"orderid"`				// A sequnce of numbers and alphabets- this will be the transaction has obtained when init-order is invoked by r retailer 
	Timestamp int64 `json:"timestamp"`			//utc timestamp of creation
	prod Description  `json:"prod"`				//description of desired product 
	Quantity_order int `json:"quantity_order"`	                       // Quantity of product needed
}

type AllOrders struct{
	Openorders []Order `json:"openorders"`
}
*/


// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// ============================================================================================================================
// Init - reset all the things
// ============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
		return nil, err
	}

return nil, nil
}


// ============================================================================================================================
// Invoke - Our entry point for Invocations
// ============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}else if function == "create_asset"{
		return t.create_asset(stub,args)
	}
	
	else if function == "query" {											//writes a value to the chaincode state
		return t.query(stub,args)
	} 
	
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation")
}

func (t *SimpleChaincode) create_asset(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if args[0] == "supplier"{
		supplierasset := &assets{User:"sarat",Quantity:args[1],Typeofasset:args[2]}
		if err != nil {
		return nil, errors.New("Failed to get marble index")
                }
		assetasbytes := stub.GetState(args[0])
		var c []assets
		json.Unmarshal(assetasbytes, &c)
		c=append(c,supplierasset)
		b,err = json.Marshal(c)
		err = stub.PutState(supplierasset.User,b)
		return nil, nil
	} else if args[0] == "retailer"{
		retailerasset := &assets{User:"ram",Quantity:args[1],Typeofasset:args[2]}
		assetasbytes := stub.GetState(args[0])
		var c []assets
		json.Unmarshal(assetasbytes, &c)
		c = append(c,retailerasset)
		b,err = json.Marshal(c)
		err = stub.PutState(retailerasset.User,b)
		return nil, nil
        }
	return nil, nil
}	
	
// ============================================================================================================================
// Query - Our entry point for Queries
// ============================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	
	
	// Handle different functions
	if function == "read" {													//read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}

// ============================================================================================================================
// Read - read a variable from chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name)									//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil													//send it onward
}



