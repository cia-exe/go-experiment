package util

import (
	"fmt"
	"strconv"
	"time"
)

// JsonTimeMs defines a timestamp encoded as epoch milliseconds/seconds in JSON
type JsonTimeMs time.Time  //milliseconds
type JsonTimeSec time.Time // seconds

//------------------------- JsonTimeMs

// MarshalJSON is used to convert the timestamp to JSON
func (t JsonTimeMs) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.UnixMilli(), 10)), nil // milliseconds
	//return []byte(strconv.FormatInt(t.Unix(), 10)), nil	 // seconds
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *JsonTimeMs) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.UnixMilli(q) // milliseconds
	//*(*time.Time)(t) = time.Unix(q, 0) // seconds
	return nil
}

func (t JsonTimeMs) UnixMilli() int64 {
	return time.Time(t).UnixMilli()
}

// UTC returns the JSON time as a time.UTC instance in UTC
func (t JsonTimeMs) UTC() time.Time {
	return time.Time(t).UTC()
}

// String returns t as a formatted string
func (t JsonTimeMs) String() string {

	// The reference time used in these layouts is the specific time stamp: 01/02 03:04:05PM '06 -0700
	return t.UTC().Format("[06/01/02 15:04:05.000]") // [21/08/30 09:48:25.1234]
	//return t.UTC().Format("[06/01/02 15:04:05]") // [21/08/30 09:48:25]
	//return fmt.Sprintf("[%v]", t.UTC()) // [2021-08-30 09:48:25 +0000 UTC]
}

//------------------------- JsonTimeSec

// MarshalJSON is used to convert the timestamp to JSON
func (t JsonTimeSec) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.Unix(), 10)), nil // seconds
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *JsonTimeSec) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0) // seconds
	return nil
}

func (t JsonTimeSec) Unix() int64 {
	return time.Time(t).Unix()
}

// UTC returns the JSON time as a time.UTC instance in UTC
func (t JsonTimeSec) UTC() time.Time {
	return time.Time(t).UTC()
}

// String returns t as a formatted string
func (t JsonTimeSec) String() string {

	// The reference time used in these layouts is the specific time stamp: 01/02 03:04:05PM '06 -0700
	return t.UTC().Format("[06/01/02 15:04:05]") // [21/08/30 09:48:25]
}

//---------------------------------------

type StrUint64 uint64 // json string to uint64

func (u *StrUint64) UnmarshalJSON(bs []byte) (err error) {

	str := string(bs) // Parse plain numbers directly.
	if bs[0] == '"' && bs[len(bs)-1] == '"' {
		// Unwrap the quotes from string numbers.
		str = string(bs[1 : len(bs)-1])
	}

	if str == "" {
		*u = StrUint64(0)
		return
	}

	x, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		fmt.Printf("[%v] str=\"%v\"\n", err, str)
		return
	}
	*u = StrUint64(x)
	return
}
