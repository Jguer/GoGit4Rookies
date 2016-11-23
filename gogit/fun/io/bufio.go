// +build OMIT

package main

// START0 OMIT

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// STOP0 OMIT

// START1 OMIT

func main() {
	// file, err := os.Open("file.txt") <- Read only
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// check if there was any error while reading
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Write to file
	if _, err = file.WriteString("HackerSchool++\n"); err != nil {
		log.Fatal(err)
	}
}

// STOP1 OMIT
