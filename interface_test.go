package main

import (
	"fmt"
	"reflect"
	"testing"
)

type iTest interface {
	fn1()
	fn2()
}

func TestEmbeddedInterface(t *testing.T) {

	var i iTest
	f1 := MyFloat1(111)
	f2 := MyFloat2{222} // for embedded struct
	//f2 := MyFloat2(222) // for type define
	v := Vertex{}

	i = f1
	fmt.Println("----------", reflect.TypeOf(i), "=", reflect.TypeOf(f1))
	i.fn1()
	i.fn2()

	i = f2
	fmt.Println("----------", reflect.TypeOf(i), "=", reflect.TypeOf(f2))
	i.fn1()
	i.fn2()

	i = &v
	fmt.Println("----------", reflect.TypeOf(i), "=", reflect.TypeOf(v))
	i.fn1()
	i.fn2()

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement iTest.
	//	a = v

}

//------------------------------
type MyFloat1 float64

func (f MyFloat1) fn1() {
	fmt.Println("f1.fn1")
}
func (f MyFloat1) fn2() {
	fmt.Println("f1.fn2")
}

//------------------------------
//type MyFloat2 MyFloat // "i = f2" cannot use f2 (variable of type MyFloat2) as iTest value in assignment: missing method fn2
type MyFloat2 struct{ MyFloat1 } // use Embedded struct to override fn1() only

func (v MyFloat2) fn1() {
	fmt.Println("f2.fn1")
}

// We don't need to Implement fn2()
// func (v MyFloat2) fn2() {
// 	fmt.Println("f2.fn2")
// }

//------------------------------
type Vertex struct {
	MyFloat1
	X, Y float64
}

func (v *Vertex) fn1() {
	v.MyFloat1.fn1() // invoke base method
	fmt.Println("v.fn1")
}

func (v *Vertex) fn2() {
	fmt.Println("v.fn2")
}

// Output:
//
// ---------- main.MyFloat1 = main.MyFloat1
// f1.fn1
// f1.fn2
// ---------- main.MyFloat2 = main.MyFloat2
// f2.fn1
// f1.fn2
// ---------- *main.Vertex = main.Vertex
// f1.fn1
// v.fn1
// v.fn2
