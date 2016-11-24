// +build OMIT

// START1 OMIT
package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Array a:", a)

	a[4] = 100
	fmt.Println("after value of a[4]:", a)
	fmt.Println("Return only one element:", a[4])

	// The builtin `len` returns the length of an array.
	fmt.Println("len:", len(a))

    // Implicit declaration is back
	b := [...]int{1, 2, 3, 4, 5}
    fmt.Println("Implict b:", b)

}

// STOP1 OMIT
