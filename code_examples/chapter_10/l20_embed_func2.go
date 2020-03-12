/*
The embedded type does not need to a pointer, Customer
does not need an Add method anymore because it uses
the method from Log, Customer can have its own String-method
and using in it the log String() method.
*/
package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name	string
    Log
}

func main() {
	c := &Customer{"Barak Obama", Log{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (c *Customer) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}

func (l *Log) String() string {
	return l.msg
}

/* Output:
Barak Obama
Log:{1 - Yes we can!
2 - After me the world will be a better place!}
*/
