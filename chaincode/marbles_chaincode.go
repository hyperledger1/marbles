/* SUPPLY CHAIN MANAGEMENT on top of Hyperledger*/

package main  


//importing the packages for using inbuilt functions
import (
	"errors"
        "bufio"
        "os"
	"fmt"
	"strconv"
	"encoding/json"
	"time"
	"strings"
        "github.com/hyperledger/fabric/core/chaincode/shim"
)

// Dont know why this has to be used..as of now
type SimpleChaincode struct {            
}

//In final phase each product will have a ID. That ID will be used as a key and this entire struct will be kept as value corresponding to that key.IF IOT is integrated
//may be by scanning bar code we can generate ID through that why also
// If the product is with supplier(As is the case initally when product is manufactured, user will be "Supplier"
//Say 10 pants are der initially(User is Supplier), market gives order for 5 pants. For 5 pants user will be set to market and for remaining 5 pants user will
//be as it is (Supplier).

type Product struct{
	Id string `json:"id"`					
	Color string `json:"color"`
	Size int `json:"size"`
                                                        // Brand string `json:"brand"` ,For now we can relax this
	User string `json:"user"`
           
}

// to facilitate order 
type Description struct{
	Color string `json:"color"`
	Size int `json:"size"`
       // Later on we should keep Brand also
}

type AnOpenOrder struct{
	Orderid string `json:"orderid"`		                 // just like practical scenario, each order placed by market will have a order ID
 	                                                        //Timestamp int64 `json:"timestamp"`			
        Want Description  `json:"want"`				//description of desired pant i.e color,size,brand etc
        Status string `json:"status"`                           // Not yet shipped, in transit, Delivered, Some items are delivered etc can be status
        Quantity int  `json:"quantity"`                         //Self explanatory
}

type currency struct{
supplycoin float32
}

type AllOrders struct{
	OpenOrders []AnOpenOrder `json:"open_orders"`           // list of all orders 
}


type Assets struct{
   User string `json:"user"`                                   // Supplier or retailer
   Prod Description   `json:"prod"`                            // the attributes of the product the entity has can be set here(for phase 1)     
   Productquantity  int  `json:"productquantity"`              // the number of products the entity has
   Coinbalance currency `json:"coinbalance"`                   //The amount of coins the asset holds, analogous to bank balance
} 



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



func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {


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
// Invoke - Our entry point for all invocations
// ============================================================================================================================


func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
fmt.Println("invoke has started working on " + function)

if function == "init" {											
		return t.Init(stub, "init", args)
	}else if function == "create_asset"{
		return t.create_asset(stub)
	}
        else if function == "write" {								       //writes a value to the ledger(leveraged using Putstate)
		return t.Write(stub, args)
	}
        else if function == "init_order" {							      //market to place an order to supplier
		return t.init_order(stub, args)
	}
        else if function == "init_logistics" {							     //supplier to trigger logistics 
		return t.init_logistics(stub, args)
	}
        else if function == "query" {								     //query for data from ledger(leveraged using Getstate)
		return t.query(stub,args)
	} 
        else if function == "sendcoins" {							     //for transfer of money between parties
		return t.sendcoins(stub,args)
	} 

// Since this is a smart contract, Once order is placed and advance is paid, automatically logistics is triggered--->Deliver....>Check and pay to supplier by market
// ....> pay to logistics by supplier
//So market will place an order, so how should logistics guy be called from a seperate function(Smart Contract is compromised here) or  from sma
}


func (t *SimpleChaincode) create_asset(stub shim.ChaincodeStubInterface) ([]byte, error) {
fmt.Println("Let's create asset, This asset creation will be done only once")

    


    
                         Supplierassets := &Assets{User:Supplier,Prod.Color:"blue",Prod.Size:32,productquantity:100,Coinbalance.supplycoins:1000}
                         supplierassetinbytes = json.Marshal(Supplierassets)           // Convering to Json format i.e bytes

                         err= stub.PutState("Supplierassets", supplierassetinbytes)   // Writing to ledger with the shown key and entire struct as value
   
                           if err != nil {
                                          fmt.Printf("Error: %s", err)
                                          return;
                           }

                        fmt.Println(string(supplierassetinbytes))                 // Printing the Contents

                       
   
  

                        Marketassets := &Assets{User:Market,Prod.Color:"blue",Prod.Size:32,productquantity:20,Coinbalance.supplycoins:1000}
                        marketassetinbytes = json.Marshal(Marketassets)           // Convering to Json format i.e bytes

                        err= stub.PutState("Marketassets", marketassetinbytes)   // Writing to ledger with the shown key and entire struct as value
   
                          if err != nil {
                                         fmt.Printf("Error: %s", err)
                                         return;
                          }

                        fmt.Println(string(marketassetinbytes))                 // Printing the Contents


                      
                       Logisticsassets:= &Assets{User:Logistics,Prod.Color:nil, Prod.Size:nil,productquantity:nil,Coinbalance.supplycoins:100}
                       logisticsassetinbytes = json.Marshal(Logisticsassets)           // Convering to Json format i.e bytes

                        err= stub.PutState("Logisticsassets", logisticsassetinbytes)   // Writing to ledger with the shown key and entire struct as value
   
                          if err != nil {
                                         fmt.Printf("Error: %s", err)
                                         return;
                          }

                        fmt.Println(string(logisticsassetinbytes))                 // Printing the Contents

     return nil, nil


}




func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)
	
	
	// Handle different functions
	if function == "read" {													//read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)						//error

	return nil, errors.New("Received unknown function query")
}


func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var name, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the variable to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name)									//get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil										       //send it onward
}




