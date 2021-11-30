package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	main()
}

func TestColorLog(t *testing.T) {

	println("error info debug warn WARN test xxx: err:")
	println("err:\033 is red font.")
	println("err\033: info debug warn WARN test xxx: err:")
	println("error\033[31mThis is red font.<INF>\033[0m")
	println("error:\033[31mThis is red font.<INF>\033[0m")
	println("err:\033[31mThis is red font.<INF>\033[0m")
	println("xxx\033[31mThis is red font.<INF>\033[0m")
	println("\033[32mThis is green font.<ERR>\033[0m")
	println("\033[33mThis is yellow font.<DBG>\033[0m")
	println("\033[34mThis is blue font.<WAR>\033[0m")
	println("\033[38mThis is the default font. \033[0m")

	println(`Colorization should work with most themes because it uses common theme token style names.
	 It also works with most instances of the output panel. Initially attempts to match common literals
	  (strings, dates, numbers, guids) and warning|info|error|server|local messages.`)

	println("\u001b[46;1mDEB:\u001b[0m @xxxxxxx@")
	println("\u001b[42;1mINF:\u001b[0m @xxxxxxx@")
	println("\u001b[41;1mERR:\u001b[0m @xxxxxxx@")

}

//--------------------------------

func TestSliceCallByRef(t *testing.T) {

	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := []int{1, 2, 3, 4, 5, 6}

	func(in []int) {
		for i, v := range in {
			in[i] = v * v
		}
	}(slice1)

	for i, v := range slice2 {
		slice2[i] = v * v
	}

	if !reflect.DeepEqual(slice1, slice2) {
		fmt.Println(slice1)
		fmt.Println(slice2)
		t.Error()
	}

	fmt.Println("call by reference OK...", slice1) //[1 4 9 16 25 36] -> slice is call by reference
}

func TestEmptyStruct(t *testing.T) {

	//-------------------------
	type class struct {
		x    int
		y    *int
		arr  []int  // nullable but String() returns []
		pArr *[]int // nullable
	}

	var c class
	fmt.Println("c1", c)                       // {0 <nil> [] <nil>}
	fmt.Println("c2", c.x, c.y, c.arr, c.pArr) // c2 0 <nil> [] <nil>

	if c.arr == nil && c.pArr == nil {
		fmt.Println("ok! Both [] and *[] are nullable.")
	}

	var d *class
	fmt.Println("d1", d) // <nil>
	//fmt.Println("d2", d.x, d.y) // error: invalid memory address or nil pointer dereference
}

func TestEmptySlice(t *testing.T) {

	{
		var arr []int
		if arr == nil {
			fmt.Println(11, arr) // []
		}
		// fmt.Println(12, arr[len(arr)-1]) //error: index out of range [-1]
		// fmt.Println(13, arr[0]) //error: index out of range [0] with length 0
	}

	{
		arr := []int{}
		if arr != nil { // NOT nil
			fmt.Println(21, arr) // []
		}
		// fmt.Println(22, arr[len(arr)-1]) //error: index out of range [-1]
		// fmt.Println(23, arr[0]) //error: index out of range [0] with length 0
	}

	{
		arr := make([]int, 0)
		if arr != nil { // NOT nil
			fmt.Println(31, arr) // []
		}
		// fmt.Println(32, arr[len(arr)-1]) //error: index out of range [-1]
		// fmt.Println(33, arr[0]) //error: index out of range [0] with length 0
	}

}

//--------------------------------

type ClassMimic struct {
	x   int
	y   int
	Str string

	fn func(int, int) int
}

func (r *ClassMimic) doPrivate() string {
	return fmt.Sprintln(r.Str, (r.x + r.y))
}

func TestClassMimic(t *testing.T) {

	var ObjS = ClassMimic{
		x:   1,
		y:   2,
		Str: "3",

		fn: func(Ma int, pay int) int {
			return Ma * pay
		},
	}

	ObjS.doPrivate()
	ObjS.fn(3, 6)
}
