// +build OMIT

// START1 OMIT
package main

import "fmt"

func main() {
	s := make([]string, 3) // HL
	fmt.Println("Slice s:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("Slice s after setting:", s)
	fmt.Println("Get one element:", s[2])
	fmt.Println("Length:", len(s))

	s = append(s, "d") // HL
	fmt.Println("Slice s after appending d:", s)
	s = append(s, "e", "f") // HL
	fmt.Println("Slice s after appending e and f:", s)

	intS := make([]int, 4) // Default value is 0
	fmt.Println("int Slice:", intS)
}

// STOP1 OMIT
