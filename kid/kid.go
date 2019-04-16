// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package kid

import (
	"encoding/hex"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"golang.org/x/crypto/sha3"
)

// GetID _
func GetID(stub shim.ChaincodeStubInterface, secure bool) (string, error) {
	creator, _ := stub.GetCreator() // error is always nil
	h := make([]byte, 20)
	sha3.ShakeSum256(h, creator)
	return hex.EncodeToString(h), nil
}
