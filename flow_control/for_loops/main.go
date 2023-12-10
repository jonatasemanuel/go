package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// i++

	// WHILE??
	j := 10 // initialization
	for j >= 0 {
		fmt.Println(j)
		j-- // post statement
	}

	// INFITE LOOp
	sum := 0
	for {
		sum++
		fmt.Println(sum)
	}
	fmt.Println(sum) // never reached
}
