package main

import "fmt"

func main() {
	// Typed
	const a float64 = 5.1

	// Untyped
	const b = 6.7

	// Constant expressions
	const c float64 = a * b
	const str = "Hello" + "Go!"

	const d = 5 > 0
	fmt.Println(d)

	// x has another type | compile error
	// const x int = 5
	// const y float64 = 2.2 * x

	// not typed works
	const x = 5
	const y = 2.2 * x // y now is float64
}
