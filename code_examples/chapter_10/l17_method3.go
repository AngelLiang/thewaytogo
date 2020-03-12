/*
it is shown that a method on an embedded struct 
can be called directly on an instance of the embedding type.
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	/*
	Embedding injects fields and methods of an existing type into another type: methods 
	associated with the anonymous field are promoted to become methods of the enclosing 
	type.
	*/
	Point
	name string
}

func main() {
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs()) // prints 5
}
