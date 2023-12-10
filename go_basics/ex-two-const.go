package main

import "fmt"

func main() {
	const secPerDay uint = 86400
	const daysYear uint = 365
	fmt.Printf("There are %d seconds in a year", secPerDay*daysYear)
}
