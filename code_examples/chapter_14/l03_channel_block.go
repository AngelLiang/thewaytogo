/*
the goroutine pump sends integers in an infinite loop on the channel. 
But because there is no receiver, the only output is the number 0
*/
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

