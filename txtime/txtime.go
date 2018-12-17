// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package txtime

import (
	"bytes"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// RFC3339NanoFixed -
// if unix time's nano seconds is 0, RFC3339Nano Format tails nano parts.
const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z"

// Time wraps go default time package
type Time struct {
	*time.Time
}

// GetTime returns the *Time converted from TxTimestamp
func GetTime(stub shim.ChaincodeStubInterface) (*Time, error) {
	ts, err := stub.GetTxTimestamp()
	if err != nil {
		return nil, err
	}
	unixTime := time.Unix(ts.GetSeconds(), int64(ts.GetNanos()))
	t := &Time{Time: &unixTime}
	return t, nil
}

// FormatRFC3339Nano returns RFC3339Nano format string that forced nano parts
func FormatRFC3339Nano(t time.Time) string {
	return t.Format(RFC3339NanoFixed)
}

// String returns RFC3339NanoFixed format string
func (t *Time) String() string {
	return t.Time.Format(RFC3339NanoFixed)
}

// MarshalJSON marshals Time as RFC3339NanoFixed format
func (t *Time) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{'"'})
	if _, err := buf.WriteString(FormatRFC3339Nano(*t.Time)); err != nil {
		return nil, err
	}
	if err := buf.WriteByte('"'); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON unmarshals RFC3339NanoFixed format bytes to Time
func (t *Time) UnmarshalJSON(data []byte) error {
	str := string(data)
	if str == "null" {
		return nil
	}
	time, err := time.Parse(`"`+RFC3339NanoFixed+`"`, str)
	if err != nil {
		return err
	}
	t.Time = &time
	return nil
}
