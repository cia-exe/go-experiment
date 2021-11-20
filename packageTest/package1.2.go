package packageTest

import "fmt"

func FunPkg1_2() {
	fmt.Println("funPkg1_2")
}

// Even if we do not import [package2b] in this file,
// it still causes a import cycle in [package2b.go]
// because packageTest.main.go import [package2b]
func Call() {
	fmt.Println("packageTest")
}
