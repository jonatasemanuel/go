/*
	Consider the following string declaration:s := "你好 Go!"

1. Convert the string to a byte slice.
2. Print out the byte slice
3. Iterate over the byte slice and print out each index and byte in the byte slice
*/
package main

import "fmt"

func main() {
	s := "你好 Go!"

	b := []byte(s)

	fmt.Printf("%#v\n", b)

	for i, v := range b {
		fmt.Printf("Index: %d Byte: %v\n", i, v)

	}
}
