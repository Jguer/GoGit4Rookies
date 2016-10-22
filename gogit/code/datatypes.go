// +build OMIT

// Single line comment
/* Multi-
line comment */

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	ToBe := true
	fmt.Printf("%T(%v)\n%T(%v)\n%T(%v)\n", ToBe, ToBe, MaxInt, MaxInt, z, z)
}
