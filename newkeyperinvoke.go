/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// NewKeyPerInvoke is allows the following transactions
//    "put", "key", val - returns "OK" on success
//    "get", "key" - returns val stored previously
type NewKeyPerInvoke struct {
}

//Init implements chaincode's Init interface
func (t *NewKeyPerInvoke) Init(stub shim.ChaincodeStubInterface) ([]byte, error) {
	return nil, nil
}

//Invoke implements chaincode's Invoke interface
func (t *NewKeyPerInvoke) Invoke(stub shim.ChaincodeStubInterface) ([]byte, error) {
        function, args := stub.GetFunctionAndParameters()

        if function != "invoke" {
                return nil, fmt.Errorf("Unknown function call")
        }

        if len(args) < 2 {
                return nil, fmt.Errorf("Incorrect number of arguments. Expecting at least 2")
        }

	//args := stub.GetArgs()
	//if len(args) < 2 {
		//return nil, fmt.Errorf("invalid number of args %d", len(args))
	//}
	f := string(args[0])
	if f == "put" {
		if len(args) < 3 {
			return nil, fmt.Errorf("invalid number of args for put %d", len(args))
		}
		err := stub.PutState(args[1], []byte(args[2]))
		if err != nil {
			return nil, err
		}
		return []byte("OK"), nil
	} else if f == "get" {
		// Get the state from the ledger
		return stub.GetState(string(args[1]))
	}
	return nil, fmt.Errorf("unknown function %s", f)
}

func main() {
	err := shim.Start(new(NewKeyPerInvoke))
	if err != nil {
		fmt.Printf("Error starting New key per invoke: %s", err)
	}
}
