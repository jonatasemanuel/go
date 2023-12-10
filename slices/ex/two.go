package main

import "fmt"

func main() {
	mySlice := []float64{1.2, 5.6}

	mySlice[1] = 6.

	a := 10.
	mySlice[0] = a

	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)

}
