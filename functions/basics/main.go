/*
- Go recommends writing function names in simple word or camelCase.
- Within the same package function names must be unique!
- One of Go's features is that funtions and methods can return multiple values
- Go doens't support function overloading
- main() and int() are predefined function names.
- Syntax:

func (receiver) name (parameters) (returns){
	// code -> function body here
}

*/

package main

import (
	"fmt"
	"math"
)

func fOne() {
	fmt.Println("This is f1() function.")
}

func fTwo(a int, b int) {
	fmt.Println("Sum:", a+b)
}

func fThree(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)

}

func fFour(a float64) float64 {
	return math.Pow(a, a)

}

func fFive(a, b int) (int, int) {
	return a + b, a * b
}

func sum(a, b int) (s int) {
	fmt.Println(s) // s -> 0

	s = a + b
	return // naked return statement
}

func main() {
	fOne()
	fTwo(5, 7)
	fThree(3, 5, 6, 4.3, 7.4, "Golang")
	p := fFour(5.1)
	fmt.Println(p)

	a, b := fFive(8, 9)

	fmt.Println(a, b)

	mySum := sum(4, 8)
	fmt.Println(mySum)
}
