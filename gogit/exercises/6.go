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
	fmt.Println("Is ", eletrao.name, "even?", eletrao.even())

	eletrao.changeCourse("MEMec", []string{"IEMec", "DMG-I", "MA-I"})
	fmt.Println(eletrao)
}

// STOP1 OMIT

// START2 OMIT
func (s student) even() bool {
	if s.istid%2 == 0 {
		return true
	}

	return false
}

func (s *student) changeCourse(course string, classes []string) {
	s.course = course
	s.classes = classes
}

// STOP2 OMIT
