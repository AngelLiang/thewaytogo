/*
Channels can be closed explicitly; however they are not like files: 
you don’t usually need to close them. Closing of a channel is only necessary 
when the receiver must be told there are no more values coming. 
Only the sender should close a channel, never the receiver.
*/
package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}
// Washington Tripoli London Beijing Tokio 
