package main

import "fmt"

func main() {
	var a = 4
	var b = 5.2

	a = int(b)
	fmt.Println(a, b)

	// var x int | A variable cannot change it-s type
	// x = "5" -> error

	// An uninitialized variable or empty variable  will get the so called ZERO VALUE
	var (
		value int
		price float64
		name  string
		done  bool
	)
	fmt.Println(value, price, name, done)

}
