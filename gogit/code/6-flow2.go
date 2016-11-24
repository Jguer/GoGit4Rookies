// +build OMIT

// START1 OMIT
package main

import "fmt"

func main() {

	i := 0
	for {
		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}

		if i < 20 {
			break
		}

		i++
	}
}

// STOP1 OMIT
