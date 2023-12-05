package main

import "fmt"

func main() {
	var age int = 23
	fmt.Println("Age: ", age)

	// Go type inference -> get the type from rigth side of equals
	var name = "Donnie"
	//fmt.Println("Name:", name)
	_ = name

	// SHORT DECLARATION | Works only in Block scope!
	s := "Learning golang!"
	fmt.Println(s)

	// Cannot use with already variables
	name = "Jony"
	name1 := "Gon"
	_ = name1

	// MULTIPLE DECLARATIONS
	car, cost := "Audi", 50000
	fmt.Println(car, cost)
	// At least one is new
	car, year := "BMW", 2018
	_ = year

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	// Multiple declaration is good for readability
	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	// MULTIPLE ASSIGNMENTS
	var i, j int
	i, j = 5, 8

	// Swapping variables
	j, i = i, j
	// _, _ = i, j
	fmt.Println(i, j)

	// Expression in short declarations are allowed
	sum := 5 + 2.3
	fmt.Println(sum)
}
