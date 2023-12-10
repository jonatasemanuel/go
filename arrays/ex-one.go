package main

import "fmt"

func main() {
	var cities [2]string
	fmt.Printf("%#v\n", cities)

	grades := [3]float64{.5, .3}
	fmt.Printf("%#v\n", grades)

	taskDone := [...]bool{}
	fmt.Printf("%#v\n", taskDone)

	cities[0], cities[1] = "Ourinhos", "Canitar"
	for i, j := range cities {
		fmt.Printf("Index: %d Values: %s Length: %d\n", i, j, len(cities))

	}
	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}
}
