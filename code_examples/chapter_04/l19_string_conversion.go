package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)	  

	an, _ = strconv.Atoi(orig)  // convert to an int
	fmt.Printf("The integer is: %d\n", an) 
	an = an + 5
	newS = strconv.Itoa(an)  // returns the decimal string representation of i
	fmt.Printf("The new string is: %s\n", newS)
}
