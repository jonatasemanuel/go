/*
	Consider the following variable declaration:x, y := 5.5, 8.8

Write a function that swaps the values of x and y.
After calling the function x will be 8.8 and y will 5.5
*/
package main

import "fmt"

func swapValues(a, b *float64) {
	*a, *b = *b, *a

}
func main() {
	x, y := 5.5, 8.8
	a, b := &x, &y
	fmt.Printf("x: %.2f | y: %.2f\n", x, y)
	swapValues(a, b)
	fmt.Printf("x: %.2f | y: %.2f\n", x, y)

}
