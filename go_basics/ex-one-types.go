package main

import "fmt"

type duration int

var hour duration

func main() {
	fmt.Println(hour)
	fmt.Printf("%T\n", hour)

	hour = 3600
	fmt.Println(hour)
}
