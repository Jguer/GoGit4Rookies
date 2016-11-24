// +build OMIT

package main

import (
	"fmt"
)

var i int = 10 // Global scope, we can remove int
const num = 3  // Constant

func main() {
	i := 42 // Overlaps global variable
	fmt.Printf("Value of i: %d\n", i)

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7/3 =", 7/3)
	fmt.Println("7%3 =", 7%3)
	fmt.Println("7*3 =", 7*3)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println("True AND false:", true && false, "\n True OR false:", true || false)
	fmt.Println("NOT true", !true)
}
