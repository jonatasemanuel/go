/* Create a function called cube() that takes a parameter of type float64 and returns the cube of that parameter (the parameter to the power of 3). */
package main

import (
	"fmt"
)

func cube(a float64) float64 {
	return a * a * a
}

func main() {
	fmt.Println(cube(3.))
}
