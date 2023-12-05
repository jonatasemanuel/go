package main

import (
	"fmt"
)

func main() {
	// Constants need a value on initialize
	const days int = 7
	const pi float64 = 3.14 // not declared
	const secondsInHour = 3600

	duration := 234 // in hours
	fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

	// RUNTIME ERROR
	// x, y := 5, 0
	// fmt.Println(x / y)

	// COMPILE TIME ERROR
	// const a = 5
	// const b = 0
	// fmt.Println(a / b)

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2 = -300
		min3 // go will repeat previous value and type. -300 in this case
	)

	fmt.Println(min1, min2, min3)

	// Constants rules
	// 1. CANNOT CHANGE A CONSTANT
	const temp int = 100
	// ~ temp = 50

	// 2.YOU CAN NOT INITIATE A CONSTANT AT RUNTIME | import some package
	// ~ const power = math.Pow(2, 3)

	// 3. YOU CAN NOT USE A VARIABLE TO INITIALIZE A CONSTANT
	// ~ t := 5
	// ~ const tc = t

	// 4. WE CAN USE THE len() FUNCTION TO INITIATE A const | len() is a built in funt
	const l1 = len("hello")
	fmt.Println(l1)
}
