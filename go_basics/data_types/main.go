package main

import "fmt"

func main() {
	// GO BUILT IN TYPES

	// INT type
	// int8: (-128) -> 127
	var b1 int8 = 127
	fmt.Printf("%T\n", b1)

	// uint16: max -> 65535
	var b2 uint16 = 65535
	fmt.Printf("%T\n", b2)

	// FLOAT type - Zero can be omitted
	var f1, f2, f3 float64 = 1.1, -0.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// RUNE type
	var r rune = 'z'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	// BOOL type
	var b bool = true
	fmt.Printf("%T\n", b)

	// STRING type: each character ocupies between 1-4 bytes.
	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	/* GO COMPOSITE TYPES */

	// Array: numered sequence of elements of a single type, fixed length
	// Slice: dynamic length

	// ARRAY type
	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)
	fmt.Printf("%d\n", numbers)

	// SLICE type
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)
	fmt.Printf("%s\n", cities)

	// MAP type | unordered group of elements of one type indexed by a set of keys
	// Similar to Python dictionary

	balances := map[string]float64{
		"USD": 2333.2,
		"EUR": 511.34,
	}
	fmt.Printf("%T\n", balances)
	fmt.Printf("%v\n", balances)

	/* STRUCT type */
	// is a sequence of named elements, called fields, each of which has name and a type
	// it's oop ?

	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you)

	pers := &you
	fmt.Printf("%s\n", you.name)
	pers.name = "Julia"

	/* POINTER type */
	// -> Stores the memory address of another variable: uninitialized pointer is nil
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	/* FUNCTION type */
	fmt.Printf("%T\n", f)
}

func f() {
}
