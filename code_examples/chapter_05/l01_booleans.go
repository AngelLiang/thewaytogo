package main

import "fmt"

func main() {
	bool1 := true
	// While ( ) around the conditions are not needed, for complex conditions they may be used to make the code clearer.
	if bool1 {
		fmt.Printf("The value is true\n")
	} else {
		fmt.Printf("The value is false\n")
	}
}
