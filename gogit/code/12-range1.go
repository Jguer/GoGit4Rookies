// +build OMIT

// START1 OMIT
package main

import (
	"fmt"
)

func main() {
	numS := []int{2, 3, 4}
	sum := 0

	for i, num := range numS { // range returns 2 values, index and value
		// It's like having i and numS[i]
		fmt.Printf("index: %d, value %d\n", i, num)
		sum += num
	}

	fmt.Println("sum:", sum)
}

// STOP1 OMIT
