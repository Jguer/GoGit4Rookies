// +build OMIT

package main

import (
	"fmt"
	"strings"
)

func main() {
	var newString string
	newString = "It's me HES"                  // We could implicit declare it
	fmt.Printf("New string: %s \n", newString) //You use %s for Strings

	fmt.Println("go" + "lang") //You can sum strings

	fmt.Println("go", "lang", 1.7) //Println receives multiple arguments

	// strings library includes many useful string options
	fmt.Println(strings.ToUpper("golang"))
}
