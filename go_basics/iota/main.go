package main

import "fmt"

func main() {
	// Zero  by default
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		c11 = iota
		c22
		c33
	)
	fmt.Println(c11, c22, c33)

	const (
		North = iota // default 0
		East
		South
		West
	)
	fmt.Println(South)

	const (
		// a = iota * 2 -> 0
		a = (iota * 2) + 1 // 1
		b                  // 3
		c                  // 5
		d                  // 7
	)
	fmt.Println(a, b, c, d)

	// Skipping values with black identifier: x = -2, y = -4, z = -5
	const (
		x = -(iota + 2) // -2
		_               // -3
		y               // -4
		z               // -5
	)

	fmt.Println(x, y, z)
}
