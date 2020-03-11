// function_parameter.go
/*
Functions can be used as parameters in another function, the passed
function can then be called within the body of that function,
that is why it is commonly called a callback.
*/
package main

import (
	"fmt"
)

func main() {
	callback(1, Add)
}

func Add(a,b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a + b)
}


func callback(y int, f func(int, int)) {
	f(y, 2)  // this becomes Add(1, 2)
}
// Output:  The sum of 1 and 2 is: 3
