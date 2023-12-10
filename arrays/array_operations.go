package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := [3]int{10, 20, 30}
	fmt.Printf("%#v\n", numbers)

	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)

	for i, v := range numbers {
		fmt.Println("Index:", i, "Value:", v)
	}

	fmt.Println(strings.Repeat("-#-", 7))
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index:", i, "Value:", numbers[i])
	}

	// Multi-dimensional
	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}
	fmt.Println(balances)

	m := [3]int{1, 2, 3}
	// another array, different alocation. Not is a copy
	n := m
	fmt.Println("n is equal to m:", n == m) // true
	m[0] = -1
	fmt.Println("n is equal to m:", n == m) // false

}
