package main

import (
	"fmt"
)

func main() {
	s1 := "I love Golango!"
	// Slicing a string returns bytes not runes
	fmt.Println(s1[2:5])

	s2 := "中文维基是世界上"
	fmt.Println(s2[0:2])

	// returning a slice of runes
	// 1st step: converting string to rune slice
	rs := []rune(s2)
	fmt.Printf("%T\n", rs)

	// 2st step: slicing the rune slic
	fmt.Println(string(rs[:3]))
}
