package main

import "fmt"

func change(a *int) *float64 {
	*a = 100

	b := 5.5
	return &b
}

func chanceVar(a int) {
	a = 66
}

func main() {
	x := 8
	p := &x // 'p' is of type pointer to int (*p -> int)

	fmt.Println("first X ->", x)
	change(p)
	fmt.Println("after X ->", x)

	fmt.Println("first X ->", x)
	chanceVar(x)
	fmt.Println("after X ->", x)
}
