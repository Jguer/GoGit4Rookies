// +build OMIT

package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	// STOP1 OMIT
}
