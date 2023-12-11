/*
- byte and rune that are aliases for **uint8** and **int32**. In Go are used to distinguish
characters from integer values
- Golang doesn't have a char data type.
- Characters our rune literals are expressed in Go by enclosing them in sigle quotes, as in
'x' or '\n', are represented using Unicode Code Points (numeric value that represents a rune literal)
- A string is a serie of bytes values. A string is a slice of bytes
and any byte slice can be encoded in a string value
- rune reoresent a single unicode character
*/
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var1, var2 := 'a', 'b' // rune -> int32 Unicode Code Points

	fmt.Printf("Type: %T, Value: %d\n", var1, var2)

	str := "tarâ"
	fmt.Println(len(str))

	fmt.Println("Byte (not rune) at position 1:", str[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	fmt.Println("\n", strings.Repeat("=-", 6))

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%c", r)
		i += size
	}

	fmt.Println("\n", strings.Repeat("=-", 6))

	for _, r := range str {
		fmt.Printf("%c", r)
	}
}
