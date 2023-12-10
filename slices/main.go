package main

import "fmt"

func main() {
	var cities []string

	fmt.Println(cities == nil)         // true
	fmt.Printf("cities %#v\n", cities) // []string(nil)
	fmt.Println(len(cities))

	// cities[0] = "London" -> runtime error, out of range
	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	friends[0] = "Gab"
	fmt.Println(friends)

	for idx, value := range numbers {
		fmt.Println(idx, value)
	}

	var n []int
	n = numbers
	fmt.Printf("%T", n)
}
