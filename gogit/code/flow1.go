// +build OMIT

package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is pair")
		} else {
			fmt.Println(i, "is not pair")
		}
	}
	// STOP1 OMIT
}
