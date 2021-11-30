package util

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestSliceToJson(t *testing.T) {

	type sClass struct {
		I1 int `json:"i1"`
		I2 int `json:"i2"`
	}

	type class struct {
		ArrPtr1 *[]int  `json:"p1"` // default = null (Slice is a pointer, Pointing to the slice is redundant!)
		ArrPtr2 *[]int  `json:"p2"` // default = null
		Arr1    []int   `json:"a1"` // default = null
		Arr2    []int   `json:"a2"` // default = null
		Cls1    sClass  `json:"c1"` // default = {"i1":0,"i2":0} (struct cannot be null!)
		Cls2    *sClass `json:"c2"` // default = null
	}

	//----------------------

	{
		c := class{}

		j, err := json.Marshal(c)
		// slice is a pointer, default is null.
		fmt.Println("******", string(j), err) //****** {"p1":null,"p2":null,"a1":null,"a2":null,"c1":{"i1":0,"i2":0},"c2":null} <nil>

		// if c.Cls1 == nil { // Error! cannot convert nil (untyped nil value) to struct
		// 	fmt.Println("ok")
		// }

		if c.Cls2 == nil {
			fmt.Println("ok")
		}

		if c.Arr1 == nil {
			fmt.Println("ok")
		}
	}

	{
		var ap1 *[]int
		var ap2 *[]int
		var a1 []int
		var a2 []int

		a1 = make([]int, 0) // empty slice
		ap1 = &a1

		c := class{
			ArrPtr1: ap1,
			ArrPtr2: ap2,
			Arr1:    a1,
			Arr2:    a2,
		}

		j, err := json.Marshal(c)
		// slice is a pointer, default is null.
		fmt.Println("******", string(j), err) //****** {"p1":[],"p2":null,"a1":[],"a2":null,"c1":{"i1":0,"i2":0},"c2":null}
	}

}

func TestTimeToJsonMs(t *testing.T) {

	tm := JsonTimeMs(time.Now())
	tm2 := JsonTimeMs{}

	j, _ := json.Marshal(tm)

	if err := json.Unmarshal(j, &tm2); err != nil {
		fmt.Println("!!!", err)
		return
	}

	if tm != tm2 {
		fmt.Println("!!! not equal !")
	}

	fmt.Printf("OK! %v->%v->%v", tm, string(j), tm2)
}

func TestTimeToJsonSec(t *testing.T) {

	tm := JsonTimeSec(time.Now())
	tm2 := JsonTimeSec{}

	j, _ := json.Marshal(tm)

	if err := json.Unmarshal(j, &tm2); err != nil {
		fmt.Println("!!!", err)
		return
	}

	if tm != tm2 {
		fmt.Println("!!! not equal !")
	}

	fmt.Printf("OK! %v->%v->%v", tm, string(j), tm2)
}
