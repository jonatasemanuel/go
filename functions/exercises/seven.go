/*
	Create a function that takes in an int value and prints out that value.

Assign the function to a variable, print out the type of the variable and then call that function through the variable name.
*/
package main

import "fmt"

func takes(a int) (r int) {
	r, _ = fmt.Println(a)
	return r
}

func main() {
	data := takes
	fmt.Printf("%T \n", data)
	data(23)
}
