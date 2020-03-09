/*
but there are no ( ) surrounding the header: for (i = 0; i < 10; i++) { } is invalid code!

the opening { has to be on the same line as the for.
*/
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
}
