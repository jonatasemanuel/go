package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		If the result of an arithmetic operation
		has more bits than can be represented in the result type. It's said to overflow
		the high order bits that do not fit are silently discard
	*/
	var x uint8 = 255
	x++ // Overflow: x now is 0
	fmt.Println(x)

	// a := uint8(255 + 1) -> compile error

	var b int8 = 127
	fmt.Printf("%d\n", b+1) // Int Overflow goes to the minimum value: -128(int8)

	b = -128
	b--
	fmt.Println(b) // Int Overflow goes to the max value: 127(int8)

	// Float number overflows to infinite
	f := float32(math.MaxFloat32)
	fmt.Println(f)

	f *= 1.2
	fmt.Println(f)
}
