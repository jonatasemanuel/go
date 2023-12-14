package main

import "fmt"

func main() {
	// &(variable)

	name := "Jowzee"
	fmt.Println(&name)

	var x int = 2
	ptr := &x

	fmt.Printf("ptr is of type %T\nValue: %v\nAdress:%p\n", ptr, ptr, &ptr)

	var po *float64 // zero value is nil - TYPE DESCRIPTION
	_ = po

	p := new(int)

	x = 100
	p = &x
	fmt.Printf("\np is of type %T\nValue: %v\nAdress: %p\n", p, p, p)

	// * is the dereferencing operator

	*p = 90 // equivalent to x = 90
	fmt.Println(x)
	fmt.Println("*p == x ->>", *p == x)

	*p = 10
	*p = *p / 2
	fmt.Println(x)

	// &value 	=> pointer
	// *pointer 	=> value
}
