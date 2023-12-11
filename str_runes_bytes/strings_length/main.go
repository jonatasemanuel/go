package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Golang"
	fmt.Println(len(s1)) // 6 bytes

	name := "Codruța"
	fmt.Println(len(name)) // 8 bytes

	n := utf8.RuneCountInString(name) // 7 runes
	m := utf8.RuneCountInString("ț")  // 1 rune
	fmt.Println(n, m)

}
