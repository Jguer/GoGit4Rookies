// +build OMIT

package main

import (
	"fmt"
)

// START4 OMIT
// START1 OMIT
func main() {
	fmt.Println(findChar("Teste um com j", 'j'))
	fmt.Println(findChar("We all love go", 'i'))
}

// STOP1 OMIT
// START3 OMIT
func findChar(toTest string, char rune) bool {
	//STOP3 OMIT
	// START2 OMIT

	for _, c := range toTest {
		if c == char {
			return true
		}
	}

	return false
}

// STOP2 OMIT
// STOP4 OMIT
