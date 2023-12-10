package main

import "fmt"

/*
- Doesn't create a new distinct type
- An alias declaration has the form: type T1 = T2
- Aliases can be used together in operations without type conversions we've seen at defined types.
*/

func main() {
	var a uint8 = 10
	var b byte
	b = a
	_ = b

	type second = uint
	var hour second = 3600
	fmt.Printf("Minutes in a hour: %d \n", hour/60)
}
