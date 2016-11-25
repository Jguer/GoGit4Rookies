// +build OMIT

package main

import (
	"fmt"
)

// START1 OMIT
func valchanger(a *int, b int) { //Pass by reference
	*a = *a + b
}
func main() {
	a := 2
	b := 3

	valchanger(&a, b)
	fmt.Println("Post-Function: ", a)
}

// STOP1 OMIT
