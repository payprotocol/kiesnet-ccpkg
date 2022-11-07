// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package contract

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/payprotocol/kiesnet-ccpkg/stringset"
	"github.com/payprotocol/kiesnet-ccpkg/txtime"
)

// Contract _
type Contract struct {
	_payload []byte
	_map     map[string]interface{}
}

// GetID implements Identifiable
func (c *Contract) GetID() string {
	return c._map["@contract"].(string)
}

// GetExpiryTime _
func (c *Contract) GetExpiryTime() (*txtime.Time, error) {
	return txtime.Parse(c._map["expiry_time"].(string))
}

// MarshalJSON _
func (c *Contract) MarshalJSON() ([]byte, error) {
	return c._payload, nil
}

// ContractCfg is configuration for 'contract' chaincode
var ContractCfg struct {
	CC string // chaincode name
}

func init() {
	if os.Getenv("DEV_CHANNEL_NAME") != "" { // dev mode
		ContractCfg.CC = "kiesnet-cc-contract"
	} else {
		ContractCfg.CC = "kiesnet-contract"
	}
}

// CreateContract invokes contract chaincode returns contract ID
func CreateContract(stub shim.ChaincodeStubInterface, doc []byte, expiry int64, signers *stringset.Set) (*Contract, error) {
	if nil == signers || signers.Size() < 2 {
		return nil, errors.New("signers must be 2+")
	}
	expb := []byte(strconv.FormatInt(expiry, 10))
	args := [][]byte{[]byte("create"), doc, expb}
	for signer := range signers.Map() {
		args = append(args, []byte(signer))
	}
	// invoke
	res := stub.InvokeChaincode(ContractCfg.CC, args, "")
	if res.GetStatus() == 200 {
		payload := res.GetPayload()
		m := make(map[string]interface{})
		err := json.Unmarshal(payload, &m)
		if err != nil {
			return nil, err
		}
		contract := &Contract{_payload: payload, _map: m}
		return contract, nil
	}
	return nil, errors.New(res.GetMessage())
}
