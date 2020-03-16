package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")  // 1
	// longWait()
	go longWait()
	// shortWait()
	go shortWait()
	fmt.Println("About to sleep in main()")  // 2
	time.Sleep(10 * 1e9) // sleep works with a Duration in nanoseconds (ns) !
	fmt.Println("At the end of main()")  // 7
}

func longWait() {
	fmt.Println("Beginning longWait()")  // 3
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")  // 6
}

func shortWait() {
	fmt.Println("Beginning shortWait()")  // 4
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")  // 5
}
/*
In main()
About to sleep in main()
Beginning longWait()
Beginning shortWait()
End of shortWait()
End of longWait()
At the end of main()
*/
