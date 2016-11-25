// +build OMIT

// START1 OMIT
package main

import (
	"fmt"
)

func main() {
	numS := []int{2, 3, 4}
	sum := 0

	for i, num := 0, numS[0]; i < len(numS); i++ {
		num = numS[i]
		fmt.Printf("index: %d, value %d\n", i, num)
		sum += num
	}

	fmt.Println("sum:", sum)
}

// STOP1 OMIT
