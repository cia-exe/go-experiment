package package3b

import (
	"fmt"

	"github.com/cia-exe/go-experiment/packageTest/package2a"
	"github.com/cia-exe/go-experiment/packageTest/package2b" // import parent ok!
)

func Call() {
	fmt.Println("package3")
	package2a.Call() // we can access sibling packages
	package2b.Call() //import cycle not allowed (cannot access the parent package)
}
