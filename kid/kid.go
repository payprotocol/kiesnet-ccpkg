// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package kid

import (
	"errors"
	"os"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// KIDCfg is configuration for invoke KID chaincode
var KIDCfg struct {
	CC   string            // chaincode name
	Args map[bool][][]byte // arguments bytes array
}

func init() {
	if os.Getenv("DEV_CHANNEL_NAME") != "" {
		KIDCfg.CC = "kiesnet-cc-id"
	} else {
		KIDCfg.CC = "kiesnet-id"
	}
	KIDCfg.Args = map[bool][][]byte{
		false: [][]byte{[]byte("kid")},              // non-secure
		true:  [][]byte{[]byte("kid"), []byte("1")}, // secure
	}
}

// GetID invokes KID chaincode and returns the kiesnet ID.
func GetID(stub shim.ChaincodeStubInterface, secure bool) (string, error) {
	res := stub.InvokeChaincode(KIDCfg.CC, KIDCfg.Args[secure], "")
	if res.GetStatus() == 200 {
		return string(res.GetPayload()), nil
	}
	return "", errors.New(res.GetMessage())
}
