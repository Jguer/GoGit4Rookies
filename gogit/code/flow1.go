// +build OMIT

package main

import (
	"fmt"
)

// START1 OMIT
func normalfor() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	return
}

// STOP1 OMIT

func main() {
	normalfor()
}
