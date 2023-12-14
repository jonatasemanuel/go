/*
	Consider the following variable declaration x := 10.10

1. Print out the address of x
2. Declare a pointer called ptr that stores the address of x.
3. Print out the type and the value of ptr.
4. Print the address of the pointer and the value of x though the pointer (use the dereferencing operator).
*/
package main

import "fmt"

func main() {
	x := 10.10
	fmt.Printf("%p\n", &x)

	ptr := &x
	fmt.Printf("prt type: %T - prt value: %v", ptr, ptr)

	fmt.Printf("pointer address: %p. | x -> %v", &ptr, *ptr)

}
