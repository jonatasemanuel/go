// INTERFACE WITH NO METHODs
// Hold values of any type

package main

import "fmt"

type emptyInterface interface{}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}

	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{4, 5, 6}
	fmt.Println(empty)
	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = "James Backster"
	fmt.Println(you.info)
	you.info = 91
	fmt.Println(you.info)
	you.info = []float64{5., 56.4, 4.}
	fmt.Println(you.info)
}
