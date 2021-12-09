package json

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

	tm := TimeMs(time.Now())
	tm2 := TimeMs{}

	j, _ := json.Marshal(&tm)

	if err := json.Unmarshal(j, &tm2); err != nil {
		fmt.Println("!!!", err)
		return
	}

	if tm != tm2 {
		fmt.Println("!!! not equal !")
	}

	fmt.Printf("OK! %v->%v->%v", &tm, string(j), &tm2)
}

func TestTimeToJsonSec(t *testing.T) {

	tm := TimeSec(time.Now())
	tm2 := TimeSec{}

	j, _ := json.Marshal(&tm)

	if err := json.Unmarshal(j, &tm2); err != nil {
		fmt.Println("!!!", err)
		return
	}

	if tm != tm2 {
		fmt.Println("!!! not equal !")
	}

	fmt.Printf("OK! %v->%v->%v", tm, string(j), tm2)
}

func TestEmbeddedStruct(t *testing.T) {

	type Inner struct {
		InnerData string `json:"jInnerData"`
	}

	{
		type Outer struct {
			Inner `json:"jInner"`
			Num   int `json:"jNum"`
		}

		o := Outer{
			//InnerData: "I'm inner",  // error: unknown field. (We cannot init Outer.InnerData!)
			Inner: Inner{InnerData: "I'm inner"},
			Num:   66,
		}

		j, _ := json.Marshal(o)
		fmt.Printf("OK1! %v->%v\n", o, string(j))
		fmt.Println("InnerData1=", o.Inner.InnerData, "==", o.InnerData) // We can access Outer.InnerData!
	}

	{
		type Outer struct {
			Inner     // `json:"jInner"`
			Num   int `json:"jNum"`
		}

		o := Outer{
			Inner: Inner{InnerData: "I'm inner"},
			Num:   77,
		}

		j, _ := json.Marshal(o)
		fmt.Printf("OK2! %v->%v\n", o, string(j))
		fmt.Println("InnerData2=", o.Inner.InnerData, "==", o.InnerData)
	}

	{ // Test 2 inner structures with the same field name.

		type Inner2 struct {
			InnerData string `json:"jInnerData"`
		}

		type Outer struct {
			Inner      // `json:"jInner"`
			Inner2     // `json:"jInner"` // Compile Warning: struct field InnerData repeats json tag "jInnerData" also at json_test.go:112
			Num    int `json:"jNum"`
		}

		o := Outer{
			Inner:  Inner{InnerData: "I'm inner"},
			Inner2: Inner2{InnerData: "I'm inner2"},
			Num:    77,
		}

		j, _ := json.Marshal(o)
		fmt.Printf("OK3! %v->%v\n", o, string(j))
		fmt.Println("InnerData3=", o.Inner.InnerData, "==", o.Inner2.InnerData)
		// fmt.Println(o.InnerData) // error: ambiguous selector o.InnerData

	}

	// output:
	// OK1! {{I'm inner} 66}->{"jInner":{"jInnerData":"I'm inner"},"jNum":66}
	// OK2! {{I'm inner} 77}->{"jInnerData":"I'm inner","jNum":77}
	// OK3! {{I'm inner} {I'm inner2} 77}->{"jNum":77}  // missing Inner and Inner2 in json string

	// InnerData1= I'm inner == I'm inner
	// InnerData2= I'm inner == I'm inner
	// InnerData3= I'm inner == I'm inner2
}

func TestEmbeddedArray(t *testing.T) {

	type Inner []int

	{
		type Outer struct {
			Inner   `json:"jInData"`
			OutData int `json:"jOutData"`
		}

		o := Outer{
			Inner:   Inner{1, 2, 3, 4, 5},
			OutData: 77,
		}

		j, _ := json.Marshal(o)
		fmt.Printf("OK1! %v->%v\n", o, string(j))
	}

	{
		type Outer struct {
			Inner       //`json:"jInData"`
			OutData int `json:"jOutData"`
		}

		o := Outer{
			Inner:   Inner{1, 2, 3, 4, 5},
			OutData: 77,
		}

		j, _ := json.Marshal(o)
		fmt.Printf("OK2! %v->%v\n", o, string(j))
	}

	// output:
	// OK1! {[1 2 3 4 5] 77}->{"jInData":[1,2,3,4,5],"jOutData":77}
	// OK2! {[1 2 3 4 5] 77}->{"Inner":[1,2,3,4,5],"jOutData":77}

}
