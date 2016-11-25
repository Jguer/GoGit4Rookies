// +build OMIT

// START1 OMIT
package main

import (
	"fmt"
)

type student struct {
	name    string
	istid   int
	course  string
	classes []string
}

func main() {
	eletrao := student{"J Guerreiro", 88888, "MEEC", []string{"FEE", "FTel", "E1", "Controlo"}}
	fmt.Println(eletrao)
	fmt.Printf("%+v \n", eletrao)
}

// STOP1 OMIT
