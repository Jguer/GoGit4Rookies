// +build OMIT

package main

import (
	"fmt"
)

func main() {
	// START1 OMIT

	for i, c := range "This is it, the last slide" {
		fmt.Println(i, c)
		fmt.Printf("%c\n", c)
	}
	// STOP1 OMIT
}
