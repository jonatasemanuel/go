package main

import "fmt"

func main() {
	numbers := []int{2, 3}
	numbers = append(numbers, 10) // can save it in a new slice

	fmt.Println(numbers)

	// append() can append more then one element
	numbers = append(numbers, 20, 40, 50)
	fmt.Println(numbers)

	n := []int{100, 200}
	numbers = append(numbers, n...)
	fmt.Println(numbers)

	/* COPY -> copies elemts into a destination slice from a source slice
	and returns the n of elemts copied, if the slices don't have the same n of elemnts
	it copies the minimum of lenght of the slices */

	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	// dst := make([]int, 2) // [10, 20]
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)
}
