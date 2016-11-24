// +build OMIT

// Single line comment
/* Multi-
line comment */

package main

import (
	"fmt"
)

func main() {
	var number int //Declare int formal
	number = 10
	fmt.Println(number) //Print line

	complexNum := 4 + 2i //Declare complex implicit
	fmt.Println(complexNum)

	var toBe bool = true      // mix
	fmt.Printf("%v \n", toBe) // Print format
	// Try using %T
}
