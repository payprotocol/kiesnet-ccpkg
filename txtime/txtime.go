// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package txtime

import (
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// RFC3339NanoFixed -
// if unix time's nano seconds is 0, RFC3339Nano Format tails nano parts.
const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z"

// GetTime returns the *Time converted from TxTimestamp
func GetTime(stub shim.ChaincodeStubInterface) (*time.Time, error) {
	ts, err := stub.GetTxTimestamp()
	if err != nil {
		return nil, err
	}
	unixTime := time.Unix(ts.GetSeconds(), int64(ts.GetNanos()))
	return &unixTime, nil
}

// FormatRFC3339Nano returns RFC3339Nano format string that forced nano parts
func FormatRFC3339Nano(t time.Time) string {
	return t.Format(RFC3339NanoFixed)
}
