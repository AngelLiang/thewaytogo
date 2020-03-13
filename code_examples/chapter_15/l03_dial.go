// dial.go.go
/*
The net.Dial function is one of the most important functions in networking.
When you Dial into a remote system the function returns a Conn interface type,
which can be used to send and receive information. The function Dial neatly abstracts 
away the network family and transport. So IPv4 or IPv6, TCP or UDP can all share a common interface.
*/
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err:= net.Dial("tcp", "192.0.32.10:80")
	checkConnection(conn, err)
	
	conn, err =net.Dial("udp", "192.0.32.10:80")
	checkConnection(conn, err)
	
	conn, err =net.Dial("tcp", "[2620:0:2d0:200::10]:80")
	checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) {
	if err!= nil {
		fmt.Printf("error %v connecting!")
		os.Exit(1)
	}
			
	fmt.Println("Connection is made with %v", conn)

}
