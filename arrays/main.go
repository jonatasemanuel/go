/*
- Go stores the elements of the array in contiguous memory locations and this way it's very efficient.
*/
package main

import (
	"fmt"
)

func main() {
	var numbers [4]int // [0 0 0 0]

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var aOne = [4]float64{}
	fmt.Printf("%#v\n", aOne)

	var aTwo = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", aTwo)

	aThree := [4]string{"Dan", "Daniel", "Juan", "Kain"}
	fmt.Printf("%#v\n", aThree)

	aFour := [4]string{"x", "y"}
	fmt.Printf("%#v\n", aFour)

	// Ellipsis operator
	aFive := [...]int{1, 2, 5, 1, -10, 66}
	fmt.Printf("%#v\n", aFive)
	fmt.Printf("The length of aFive is %d\n", len(aFive))

	aSix := [...]int{1,
		2,
		5,
		1,
		-10,
		66, // this comma is mandatory
	}
	fmt.Println(aSix)
}
