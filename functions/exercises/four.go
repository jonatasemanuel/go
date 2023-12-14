/* Create a function with the identifier sum that takes in a variadic parameter of type int and returns the sum of all values of type int passed in. */
package main

import "fmt"

func sum(a ...int) (s int) {
	for _, v := range a {
		s += v
	}
	return s
}

func main() {
	fmt.Println(sum(3, 6, 9))
}
