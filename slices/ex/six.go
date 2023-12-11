/*
	Consider the following slice declaration:

Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.
*/
package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	bF := []string{}
	bF = append(bF, friends...)
	/*
		exercice 8
			bF := make([]string, len(friends))
			copy(bF, friends)
	*/
	bF[1] = "Jane"
	fmt.Printf("Friends %#v\nNow %#v\n", friends, bF)
}
