package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "Go is cool!"
	x := s1[0]
	fmt.Println(x)
	fmt.Printf("%c\n", x)

	// s1[0] = 'x'

	// printing the number of runes in the string
	count := utf8.RuneCountInString(s1)
	fmt.Println(count)
}
