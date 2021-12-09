package performance

// call by value vs call by reference test
import (
	"testing"
)

//-------------------------------------------------------
type class struct {
	x0 int
	x1 int
	x2 int
	x3 int
	x4 int
	x5 int
	x6 int
	x7 int
	x8 int
	x9 int
}

func newClass(seed int) *class {
	x := seed
	return &class{
		x, x + 1, x + 2, x + 3, x + 4, x + 5, x + 6, x + 7, x + 8, x + 9,
	}
}

func (t *class) CallByRef() int {
	x := 0
	for i := 0; i < 33; i++ { // Indirect access is slower than direct access
		x += i + t.x0 + t.x1 + t.x2 + t.x3 + t.x4 + t.x5 + t.x6 + t.x7 + t.x8 + t.x9
	}
	return x
}

func (t class) CallByVal() int {
	x := 0
	for i := 0; i < 33; i++ {
		x += i + t.x0 + t.x1 + t.x2 + t.x3 + t.x4 + t.x5 + t.x6 + t.x7 + t.x8 + t.x9
	}
	return x
}

//-------------------------------------------------------
type classArrSmall struct {
	vals [99]int
}

func newClassArrSmall() *classArrSmall {
	c := classArrSmall{}
	for i := 0; i < len(c.vals); i++ {
		c.vals[i] = i + 12345
	}
	return &c
}

func (t classArrSmall) CallByVal() int {
	x := 12345
	for v := range t.vals {
		x += v
	}
	return x
}

func (t *classArrSmall) CallByRef() int {
	x := 12345
	for v := range t.vals {
		x += v
	}
	return x
}

//-------------------------------------------------------
type classArrLarge struct {
	vals [99999]int
}

func newClassArrLarge() *classArrLarge {
	c := classArrLarge{}
	for i := 0; i < len(c.vals); i++ {
		c.vals[i] = i + 12345
	}
	return &c
}

func (t classArrLarge) CallByVal() int {
	x := 12345
	for v := range t.vals {
		x += v
	}
	return x
}

func (t *classArrLarge) CallByRef() int {
	x := 12345
	for v := range t.vals {
		x += v
	}
	return x
}

//-------------------------------------------------------
type classSlice struct {
	vals []int
}

func newClassSlice(size int) *classSlice {
	c := classSlice{make([]int, size)}
	for i := 0; i < len(c.vals); i++ {
		c.vals[i] = i + 12345
	}
	return &c
}

func (t classSlice) CallByVal() int {
	x := 12345
	for v := range t.vals {
		x += v
	}
	return x
}

func (t *classSlice) CallByRef() int {
	x := 12345
	for v := range t.vals {
		x += v
	}
	return x
}

//-------------------------------------------------------

func BenchmarkCallByRef_arrSmall(b *testing.B) {
	c := newClassArrSmall()
	acc := 0
	for i := 0; i < b.N; i++ {
		acc += c.CallByRef()
	}
	_ = acc
}

func BenchmarkCallByVal_arrSmall(b *testing.B) {
	c := newClassArrSmall()
	acc := 0
	for i := 0; i < b.N; i++ {
		acc += c.CallByVal()
	}
	_ = acc
}

//--------------------
func BenchmarkCallByRef_arrLarge(b *testing.B) {
	c := newClassArrLarge()
	acc := 0
	for i := 0; i < b.N; i++ {
		acc += c.CallByRef()
	}
	_ = acc
}

func BenchmarkCallByVal_arrLarge(b *testing.B) {
	c := newClassArrLarge()
	acc := 0
	for i := 0; i < b.N; i++ {
		acc += c.CallByVal()
	}
	_ = acc
}

//--------------------
func BenchmarkCallByRef_sliceSmall(b *testing.B) {
	c := newClassSlice(99)
	acc := 0
	for i := 0; i < b.N; i++ {
		acc += c.CallByRef()
	}
	_ = acc
}

func BenchmarkCallByVal_sliceSmall(b *testing.B) {
	c := newClassSlice(99)
	acc := 0
	for i := 0; i < b.N; i++ {
		acc += c.CallByVal()
	}
	_ = acc
}

//--------------------
func BenchmarkCallByRef_sliceLarge(b *testing.B) {
	c := newClassSlice(99999)
	acc := 0
	for i := 0; i < b.N; i++ {
		acc += c.CallByRef()
	}
	_ = acc
}

func BenchmarkCallByVal_sliceLarge(b *testing.B) {
	c := newClassSlice(99999)
	acc := 0
	for i := 0; i < b.N; i++ {
		acc += c.CallByVal()
	}
	_ = acc
}

//--------------------

func BenchmarkCallByRef_class(b *testing.B) {
	c := newClass(12345)
	d := newClass(67890)
	acc := 0
	for i := 0; i < b.N; i++ {
		acc += c.CallByRef() + d.CallByRef()
	}
	_ = acc
}

func BenchmarkCallByVal_class(b *testing.B) {
	c := newClass(12345)
	d := newClass(67890)
	acc := 0
	for i := 0; i < b.N; i++ {
		acc += c.CallByVal() + d.CallByVal()
	}
	_ = acc
}

//==========================
// Conclusion & Assumption
//
// Call by reference is faster than call by value, especially for large arrays, because it does not need to copy many elements.
// The difference for Slice is very small, because Slice itself is a pointer(?) and does not copy every elements.
// Indirect access is slower than direct access to Struct. But there is no difference for Array/Slice, because they are native indirect access(?).
//

//-------------- test results -------------
//
// go version go1.17 windows/amd64
// cpu: Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz

// BenchmarkCallByRef_arrSmall
// 28982076	        39.41 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_arrSmall
// 23618976	        51.51 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByRef_arrLarge
//    38797	     30425 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_arrLarge
//    20769	     56693 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByRef_sliceSmall
// 29594188	        39.48 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_sliceSmall
// 28561296	        40.93 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByRef_sliceLarge
//    39279	     30275 ns/op	      20 B/op	       0 allocs/op
// BenchmarkCallByVal_sliceLarge
//    38937	     31137 ns/op	      20 B/op	       0 allocs/op
// BenchmarkCallByRef_class
// 38003185	        31.15 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_class
// 53444734	        22.11 ns/op	       0 B/op	       0 allocs/op
// PASS

// BenchmarkCallByRef_arrSmall
// 29089075	        39.41 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_arrSmall
// 23503004	        51.14 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByRef_arrLarge
//    38352	     30417 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_arrLarge
//    20660	     57861 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByRef_sliceSmall
// 27734250	        39.85 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_sliceSmall
// 29180040	        40.73 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByRef_sliceLarge
//    39825	     30543 ns/op	      20 B/op	       0 allocs/op
// BenchmarkCallByVal_sliceLarge
//    38509	     31018 ns/op	      20 B/op	       0 allocs/op
// BenchmarkCallByRef_class
// 37116309	        31.25 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_class
// 55933624	        21.86 ns/op	       0 B/op	       0 allocs/op
// PASS

// BenchmarkCallByRef_arrSmall
// 29840947	        39.13 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_arrSmall
// 21735744	        50.61 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByRef_arrLarge
//    39696	     30984 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_arrLarge
//    20436	     62950 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByRef_sliceSmall
// 30006151	        43.19 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_sliceSmall
// 25828891	        41.57 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByRef_sliceLarge
//    36920	     30827 ns/op	      21 B/op	       0 allocs/op
// BenchmarkCallByVal_sliceLarge
//    37969	     30884 ns/op	      21 B/op	       0 allocs/op
// BenchmarkCallByRef_class
// 37427250	        31.31 ns/op	       0 B/op	       0 allocs/op
// BenchmarkCallByVal_class
// 50375929	        21.77 ns/op	       0 B/op	       0 allocs/op
// PASS
