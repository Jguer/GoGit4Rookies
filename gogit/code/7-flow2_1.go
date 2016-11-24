// +build OMIT
// START1 OMIT

package main

import (
	"fmt"
)

func main() {

	i := 0
	for i < 20 {

		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}

		i++

	}
}

// STOP1 OMIT
