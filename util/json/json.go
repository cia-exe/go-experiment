package json

import (
	"fmt"
	"strconv"
	"time"
)

// TimeMs defines a timestamp encoded as epoch milliseconds/seconds in JSON
type TimeMs time.Time  //milliseconds
type TimeSec time.Time // seconds

//------------------------- JsonTimeMs

// MarshalJSON is used to convert the timestamp to JSON
func (t *TimeMs) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.UnixMilli(), 10)), nil // milliseconds
	//return []byte(strconv.FormatInt(t.Unix(), 10)), nil	 // seconds
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *TimeMs) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.UnixMilli(q) // milliseconds
	//*(*time.Time)(t) = time.Unix(q, 0) // seconds
	return nil
}

func (t *TimeMs) UnixMilli() int64 {
	return time.Time(*t).UnixMilli()
}

// UTC returns the JSON time as a time.UTC instance in UTC
func (t *TimeMs) UTC() time.Time {
	return time.Time(*t).UTC()
}

// String returns t as a formatted string
func (t TimeMs) String() string {

	// The reference time used in these layouts is the specific time stamp: 01/02 03:04:05PM '06 -0700
	return t.UTC().Format("[06/01/02 15:04:05.000]") // [21/08/30 09:48:25.1234]
	//return t.UTC().Format("[06/01/02 15:04:05]") // [21/08/30 09:48:25]
	//return fmt.Sprintf("[%v]", t.UTC()) // [2021-08-30 09:48:25 +0000 UTC]
}

func (t TimeMs) Equal(u TimeMs) bool {
	return t.UnixMilli() == u.UnixMilli() // Precision in milliseconds
	//return time.Time(t).Equal(time.Time(u)) // t.nsec() != u.nsec()
}

//------------------------- JsonTimeSec

// MarshalJSON is used to convert the timestamp to JSON
func (t *TimeSec) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.Unix(), 10)), nil // seconds
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *TimeSec) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0) // seconds
	return nil
}

func (t *TimeSec) Unix() int64 {
	return time.Time(*t).Unix()
}

// UTC returns the JSON time as a time.UTC instance in UTC
func (t *TimeSec) UTC() time.Time {
	return time.Time(*t).UTC()
}

// String returns t as a formatted string
func (t TimeSec) String() string {

	// The reference time used in these layouts is the specific time stamp: 01/02 03:04:05PM '06 -0700
	return t.UTC().Format("[06/01/02 15:04:05]") // [21/08/30 09:48:25]
}

func (t TimeSec) Equal(u TimeSec) bool {
	return t.Unix() == u.Unix() // Precision in seconds
	//return time.Time(t).Equal(time.Time(u)) // t.nsec() != u.nsec()
}

//---------------------------------------

// Define a Int that can accept numeric string with quotes (but it is not encoded with quotes).
type IntStr int64 // json string to uint64

// If we do not define it, json.Marshal() will call the method of base type through reflect, which will return a numeric string without quotes too.
// func (i IntStr) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(int64(i)) // call the method of bese struct
// 	//return []byte(strconv.FormatInt(int64(i), 10)), nil // int to str
// }

func (i *IntStr) UnmarshalJSON(b []byte) (err error) {

	str := string(b) // Parse plain numbers directly.
	if b[0] == '"' && b[len(b)-1] == '"' {
		// Unwrap the quotes from string numbers.
		str = string(b[1 : len(b)-1])
	}

	if str == "" {
		*i = IntStr(0)
		return
	}

	x, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("[%v] str=\"%v\"\n", err, str)
		return
	}
	*i = IntStr(x)
	return
}
