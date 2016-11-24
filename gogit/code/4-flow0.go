// +build OMIT

package main

// START1 OMIT
import "fmt"
import "strings" // importing line by line

func main() {
	coolString := "Go Workshop"
	if coolString[0] == 'G' {
		fmt.Println("Correct. The first letter is G")
	} else {
		fmt.Printf("False. The first letter is %c", coolString[0])
	}

	if strings.Contains("Hacker Entertainment System", "H") {
		fmt.Println("The string contains a capital 'H'")
	}

	if num := 9; num < 0 { // Change num
		fmt.Println(num, "is negative")
	} else if 0 <= num && num < 10 { // Multiple conditions
		fmt.Println(num, "is between 1 and 10")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// STOP1 OMIT
