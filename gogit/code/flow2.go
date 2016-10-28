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
			fmt.Println(i, "is pair")
		} else {
			fmt.Println(i, "is not pair")
		}
		i++
        if i == 20 {
            break;
        }
	}
	// STOP1 OMIT
}
