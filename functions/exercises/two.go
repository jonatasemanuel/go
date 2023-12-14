/*
	Create a Go program with a function called f1() that takes a parameter of type uint and returns 2 values:

a) the factorial of n
b) the sum of all integer numbers greater than zero (>0) and less than or equal to n (<=n)
*/
package main

import "fmt"

func f1(n uint) (int, int) {
	fact := 1
	sum := 0

	for i := 1; i <= int(n); i++ {
		fact *= i
		sum += i
	}
	return fact, sum
}

func main() {
	fat, sum := f1(5)
	fmt.Println("!:", fat, "Sum:", sum)
}
