package main

import "fmt"

func main() {
	var i int = 3
	var f float64 = 3.2

	j, k := float64(i), int(f)

	fmt.Printf("j types %T value: %v k types %T value: %v\n", j, j, k, k)

}
