// +build OMIT

// START1 OMIT
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	byteS, err := ioutil.ReadFile("file.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(byteS) // print the content as 'bytes'

	str := string(byteS) // convert content to a 'string'
	fmt.Println(str)     // print the content as a 'string'

	// Write
	str = str + "HackerSchool++\n"
	ioutil.WriteFile("file.txt", []byte(str), 0666)

}

// STOP1 OMIT
