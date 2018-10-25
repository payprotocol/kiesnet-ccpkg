// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package txtime

import (
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// GetTime returns the *Time converted from TxTimestamp
func GetTime(stub shim.ChaincodeStubInterface) (*time.Time, error) {
	ts, err := stub.GetTxTimestamp()
	if err != nil {
		return nil, err
	}
	unixTime := time.Unix(ts.GetSeconds(), int64(ts.GetNanos()))
	return &unixTime, nil
}
