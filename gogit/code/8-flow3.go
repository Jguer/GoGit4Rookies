// +build OMIT

// START1 OMIT
package main

import "fmt"

func main() {

	i := 2
	fmt.Print("write ", i, " as ") //Same as println but without \n in the end
	switch i {
	case 1: // if i == 1
		fmt.Println("one")
	case 2: // if i == 2
		fmt.Println("two")
	case 3, 4, 5, 6:
		fmt.Println("something")
	default:
		fmt.Println("don't care")
	}
}

// STOP1 OMIT
