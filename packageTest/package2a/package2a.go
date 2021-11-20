package package2a

import (
	"fmt"
	//"github.com/cia-exe/go-experiment/packageTest/package2b/package3b" // import cycle not allowed
)

func Call() {
	fmt.Println("package2a")
}

func CallPkg3() {
	fmt.Println("package2a.CallPkg3")
	//package3b.Call() // import cycle not allowed even if it is sibling.(package3b imports package2a already)
}
