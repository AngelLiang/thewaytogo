/*
If key1 does not exist in the map, val1 becomes the zero-value for the value’s type.

判断一个 key 是否存在于 map

	if _, ok := map1[key1]; ok {
		// ...
	}

*/
package main

import "fmt"

func main() {
	var value int
	var isPresent bool

	map1 := make(map[string]int)
	map1["New Delhi"] = 55
	map1["Bejing"] = 20
	map1["Washington"] = 25

	value, isPresent = map1["Bejing"]
	if isPresent {
		fmt.Printf("The value of \"Bejing\" in map1 is: %d\n", value)  // The value of "Bejing" in map1 is: 20
	} else {
		fmt.Println("map1 does not contain Bejing")
	}

	value, isPresent = map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)  // Is "Paris" in map1 ?: false
	fmt.Printf("Value is: %d\n", value)  // Value is: 0

	// delete an item:
	delete(map1, "Washington")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Washington")
	}
}
