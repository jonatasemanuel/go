/*
 1. Declare a variable called empty of type empty interface. Print out its type.

2. Assign an int value to the variable called empty. Print out its type.
3. Assign a float64 value to empty. Print out its type.
4. Assign an int slice value to empty. Print out its type.
5. Add a new int value to the slice (empty variable).
6. Print out the slice (empty variable).
*/
package main

import "fmt"

func main() {
	var empty interface{}
	fmt.Printf("%T\n", empty)

	empty = 23
	fmt.Printf("%T\n", empty)

	empty = []int{}
	fmt.Printf("%T\n", empty)

	empty = []int{23, 2, 4}

	fmt.Printf("%#v\n", empty)

}
