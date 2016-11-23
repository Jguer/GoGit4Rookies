package main

import "fmt"
import "math"

// START1 OMIT
// Here's a basic interface for geometric shapes.
type geometry interface {
	area() float64
	perim() float64
}

// STOP1 OMIT

// START2 OMIT
// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// STOP2 OMIT

// START3 OMIT
// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for `circle`s.
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// STOP3 OMIT
// START4 OMIT

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	fmt.Println("Printing Rectangle measures")
	measure(r)

	fmt.Println("Printing Circle measures")
	measure(c)
}

// STOP4 OMIT
