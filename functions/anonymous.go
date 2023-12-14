package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!")

	a := increment(10)
	fmt.Printf("%T ", a)
	a()              // 11
	fmt.Println(a()) // 12
	fmt.Println(a()) // 13
	fmt.Println(a()) // 14
}