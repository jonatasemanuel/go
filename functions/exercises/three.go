/*
	Write a function called myFunc() that takes exactly one argument which is an int number written between double quotes (this is in fact a string). If the argument is integer 'n', the function should return the result of n + nn + nnn

Example: myFunc('5') returns 5 + 55 + 555 which is 615 and myFunc('9') returns 9 + 99 + 999 which is 1107
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func myFunc(a string) int {

	r, err := strconv.Atoi(a)
	if err != nil {
		fmt.Printf("Cannot convert %q to int.", a)
		os.Exit(1) //exit the program
	}
	rr, _ := strconv.Atoi(a + a)
	rrr, _ := strconv.Atoi(a + a + a)

	return r + rr + rrr
}

func main() {
	fmt.Println(myFunc("5"))
}
