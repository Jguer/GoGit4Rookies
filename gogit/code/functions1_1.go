// +build OMIT

package main

import (
	"fmt"
)

// START1 OMIT
func main() {
	funky := func(a, b int) int {
		return (a + b)
	} // It's a func inside a func

	fmt.Println("Sum 3 and 5: ", funky(3, 5))

	fmt.Println("Sum 3 and 5: ",
		func(a, b int) int {
			return (a + b)
		}(3, 5)) // Called with args 3 and 5

}

// STOP1 OMIT
