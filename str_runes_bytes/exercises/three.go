/*
	Consider the following string declaration: s1 := "țară means country in Romanian"

1. Iterate over the string and print out byte by byte

2. Iterate  over the string and print out rune by rune and the byte index where the rune starts in the string
*/
package main

import "fmt"

func main() {
	s1 := "țară means country in Romanian"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%v ", s1[i])
	}

	for i, v := range s1 {
		fmt.Printf("Index: %v Rune: %c\n", i, v)
	}
}
