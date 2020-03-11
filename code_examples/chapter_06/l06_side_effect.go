// side_effect.go
/*
Passing a pointer to a function not only conserves memory
because no copy of the value is made.
It has also as side-effect that the variable or object can be changed
inside the function, so that the changed object doesn’t have to be
returned back from the function. See this in the following little program,
where reply, a pointer to an integer, is being changed in the function itself
*/
package main

import (
	"fmt"
)

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
	fmt.Println("Multiply:", n) // Multiply: 50
}
