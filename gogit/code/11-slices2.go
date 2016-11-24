// +build OMIT

package main

import "fmt"

func main() {
	// START1 OMIT
	newSlice := []string{"Frase 1", "Frase 2", "Frase 3", "Frase 4", "Frase 5", "Frase 6"}
	fmt.Println("NewSlice:", newSlice)
	fmt.Println("NewSlice size:", len(newSlice))

	boundedSlice := newSlice[2:5] // HL
	fmt.Println("sl1:", boundedSlice)

	lowerSlice := newSlice[:5] // HL
	fmt.Println("sl2:", lowerSlice)

	UpperSlice := newSlice[2:] // HL
	fmt.Println("sl3:", UpperSlice)
	// STOP1 OMIT
}
