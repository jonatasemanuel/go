package main

import "fmt"

func main() {
	var n []int
	fmt.Println(n == nil) // TRUE

	m := []int{}          // have a initialize
	fmt.Println(m == nil) // FALSE

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	/*
		fmt.Println(a == b)
		// SLICE CAN ONLY BE COMPARED TO NIL
		To compare with another slice we need a for loop
	*/

	var eq bool = true

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}

	if len(a) != len(b) {
		eq = false
	}

	if eq {
		fmt.Println("A and B are equals")
	} else {
		fmt.Println("Not equals")
	}

}
