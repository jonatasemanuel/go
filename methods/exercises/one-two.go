/*
	Coding Exercise #1

1. Create a new type called money. Its underlying type is float64.
2. Create a method called print that prints out the money value with exact 2 decimal points.
Are you stuck? Do you want to see the solution for this exercise? Click here.
Coding Exercise #2
Consider the money type declared at exercise #1. Create a new method for the money type called printStr that returns the money value as a string with 2 decimal points.
*/
package main

import "fmt"

type money float64

func (c money) print() {
	fmt.Printf("%.2f\n", c)
}

func (s money) printStr() string {
	str := fmt.Sprintf("%.2f", s)
	return str
}

func main() {
	var myMoney money = 14.3
	myMoney.print()
	str := myMoney.printStr()
	fmt.Printf("Type: %T | Value: %v\n", str, str)
}
