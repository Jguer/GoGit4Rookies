// +build OMIT

package main

// START1 OMIT
import (
	"fmt"
)

func main() {
	fruitS := []string{"banana", "kiwi", "apple", "pineapple"}
	fmt.Printf("%+v\n", fruitS)

	for i := range fruitS {
		// You'll make mistakes here if you're not careful
		// Use _ , i
		fmt.Println("Got:", i)
	}
}

// STOP1 OMIT
