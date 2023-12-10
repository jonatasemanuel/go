/*
Create a Go program that reads some numbers from the command line and then calculates the sum and the product of all the numbers given at the command line.
The user should give between 2 and 10 numbers.
Notes:
- the program should be run in a terminal (go run main.go) not in Go Playground
- example:
go run main.go 3 2 5
Expected output: Sum: 10, Product: 30
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Entry need between 2 and 10")
	} else {
		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				fmt.Println("Error type")
			}
			product *= num
			sum += num

		}
		fmt.Printf("Sum: %.2f Product: %.2f\n", sum, product)
	}
}
