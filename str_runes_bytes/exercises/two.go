/*
 1. Declare a rune called r that stores the non-ascii letter ă

2. Print out the type of r
3. Declare 2 strings that contains the values ma and m
4. Concatenate the strings and the rune in a new string (the new string will hold the value mamă ).
BTW: mamă means mother in Romanian.
Note: You should convert rune to string to contatenate it to another string.
*/
package main

import "fmt"

func main() {
	r := 'ă'
	fmt.Printf("%T\n", r)

	first, secon := "ma", 'm'
	newString := first + string(secon) + string(r)

	fmt.Println(newString)
}
