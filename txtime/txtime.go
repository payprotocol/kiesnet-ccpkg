// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package txtime

import (
	"bytes"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)

// RFC3339NanoFixed -
// if unix time's nano seconds is 0, RFC3339Nano Format tails nano parts.
const RFC3339NanoFixed = "2006-01-02T15:04:05.000000000Z"

// Time wraps go default time package
type Time struct {
	time.Time
}

// GetTime returns the *Time converted from TxTimestamp
func GetTime(stub shim.ChaincodeStubInterface) (*Time, error) {
	ts, err := stub.GetTxTimestamp()
	if err != nil {
		return nil, err
	}
	t := Unix(ts.GetSeconds(), int64(ts.GetNanos()))
	// check fogery	±5 minutes
	diff := time.Now().Sub(t.Time).Minutes()
	if diff > 5 || diff < -5 {
		return nil, errors.Errorf("txtime is out of bound [%.0f]", diff)
	}
	return t, nil
}

// New _
func New(t time.Time) *Time {
	return &Time{t}
}

// Parse parses a RFC3339 formatted string and returns the time value it represents.
func Parse(value string) (*Time, error) {
	t, err := time.Parse(RFC3339NanoFixed, value)
	if err != nil {
		return nil, err
	}
	return New(t), nil
}

// Unix returns the local *Time corresponding to the given Unix time, sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// https://godoc.org/time#Unix
func Unix(sec int64, nsec int64) *Time {
	return &Time{time.Unix(sec, nsec)}
}

// Cmp - before returns -1, after returns 1, equal returns 0
func (t *Time) Cmp(c *Time) int {
	if t.Time.Before(c.Time) {
		return -1
	}
	if t.Time.After(c.Time) {
		return 1
	}
	return 0
}

// String returns RFC3339NanoFixed format string
func (t *Time) String() string {
	return t.Time.Format(RFC3339NanoFixed)
}

// MarshalJSON marshals Time as RFC3339NanoFixed format
func (t *Time) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{'"'})
	if _, err := buf.WriteString(t.String()); err != nil {
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
	time, err := time.Parse(RFC3339NanoFixed, strings.Trim(str, `"`))
	if err != nil {
		return err
	}
	t.Time = time
	return nil
}
