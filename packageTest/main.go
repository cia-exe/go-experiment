package packageTest

import (
	"github.com/cia-exe/go-experiment/packageTest/package2a"
	"github.com/cia-exe/go-experiment/packageTest/package2b"
	"github.com/cia-exe/go-experiment/packageTest/package2b/package3b"
)

func main() {

	FunPkg1_1()
	FunPkg1_2()

	package2a.Call()
	package2b.Call()

	package3b.Call()
}
