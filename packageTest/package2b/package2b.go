package package2b

import (
	"fmt"

	//"github.com/cia-exe/go-experiment/packageTest" //import cycle not allowed
	"github.com/cia-exe/go-experiment/packageTest/package2a"
)

func Call() {
	fmt.Println("package2b")
	package2a.Call() // we can access sibling packages
	//packageTest.Call() //import cycle not allowed (cannot access the parent package that imports child)
}
