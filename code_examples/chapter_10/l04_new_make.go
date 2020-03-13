// annoy1.go
package main

type Foo map[string]string

// 结构体
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// OK:
	// 可以对 struct new
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1
	
	// not OK:
	// 不能对 struct 进行 make
	z := make(Bar) // compile error: cannot make type Bar
	z.thingOne = "hello"
	z.thingTwo = 1

	// OK:
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// not OK:
	u := new(Foo)
	(*u)["x"] = "goodbye" // !! panic !!: runtime error: assignment to entry in nil map
	(*u)["y"] = "world"
}
