package main

import "fmt"

func main() {
	birthYear, currentYear := 1999, 2023

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%v ", i)
		i++
	}
}
