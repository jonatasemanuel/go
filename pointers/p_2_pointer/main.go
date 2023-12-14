package main

import "fmt"

func main() {
	a := 5.5
	p1 := &a

	pp1 := &p1

	fmt.Printf("Value of p1: %v\nAdress p1: %v\n", p1, &p1)     // A address
	fmt.Printf("Value of pp1: %v\nAdress pp1: %v\n", pp1, &pp1) // p1 address

	fmt.Printf("*p1 is %v\n*pp1 is %v\n", *p1, *pp1)

	fmt.Printf("**pp1 is %v\n", **pp1)
	**pp1 = 10
	**pp1++
	fmt.Println("\n(a) now:", a)

	//** COMPRING POINTERS**//
	// - The zero value for a pointer of any type is nil.
	// - It's true when the pointer is the same varible point

	var p2 *int
	fmt.Printf("%#v\n", p2) // (*int)(nil)

	y := 5
	p2 = &y // y adress

	z := 5
	p3 := &z // z adress

	fmt.Println(p2 == p3) // false

	p4 := &y
	fmt.Println(p2 == p4) // true
}
