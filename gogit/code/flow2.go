// +build OMIT

package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	i := 0
	for {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
		i++
		if i == 20 {
			break
		}
	}
	// STOP1 OMIT
}
