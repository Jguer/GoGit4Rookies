// +build OMIT

package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"] // HL
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2") // HL
	fmt.Println("map:", m)

	_, err := m["k2"] // HL
	fmt.Println("found k2:", err)

	n := map[string]string{"Social": "2.65EUR", "Bar de Civil": "3EUR"}
	fmt.Println("map:", n)
	// STOP1 OMIT
}
