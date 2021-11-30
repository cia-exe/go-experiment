package main

import (
	"fmt"
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

type ClassX struct {
	x   int
	y   int
	Str string

	fn func(int, int) int
}

func (r *ClassX) doPrivate() string {
	return fmt.Sprintln(r.Str, (r.x + r.y))
}

var ObjS = ClassX{
	x:   1,
	y:   2,
	Str: "3",

	fn: func(Ma int, pay int) int {
		return Ma * pay
	},
}

func TestClassX(t *testing.T) {

	ObjS.doPrivate()
	ObjS.fn(3, 6)
}
