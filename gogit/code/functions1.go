// +build OMIT

package main

import (
	"fmt"
)

// START1 OMIT
func plus(a int, b int) int { // 2 args, 1 return
	return a + b
}

func morePlus(a, b, c int) (int, error) { // 2 args, 2 returns
	return a + b + c, nil
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	fmt.Println("1+2 =", plus(1, 2))

	res, err := morePlus(1, 2, 3)
	if err != nil {
		fmt.Println("Something bad happened")
	}
	fmt.Println("1+2+3 =", res)
}

// STOP1 OMIT
