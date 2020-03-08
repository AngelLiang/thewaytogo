/*
- After import: declare constants, variables and the types
- Then comes the init() function if there is any: this is a special function that every package  can contain and that is executed first.
- Then comes the main() function (only in the package main)
- Then come the rest of the functions, the methods on the types first; or the functions in order as they are called from main() onwards; or the methods and functions alphabetically if the number of functions is high.
*/
package main

import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() {
	// initialization of package
}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	//...
}

func Func1() { // exported function Func1
	//...
}
