/*
The defer keyword allows us to postpone the execution of a statement
or a function until the end of the enclosing (calling) function:
it executes something (a function or an expression) when
the enclosing function returns (after every return and even when 
an error occurred in the midst of executing the function, not only
a return at the end of the function), but before the } (Why after?
Because the return statement itself can be an expression which does something instead of only giving back 1 or more variables).
*/
package main

import "fmt"

func main() {
	Function1()
}

func Function1() {
	fmt.Printf("In Function1 at the top\n")
	defer Function2()  // 在函数返回前执行
	fmt.Printf("In Function1 at the bottom!\n")
}

func Function2() {
	fmt.Printf("Function2: Deferred until the end of the calling function!")
}
/* output:
In Function1 at the top
In Function1 at the bottom!
Function2: Deferred until the end of the calling function!
*/
