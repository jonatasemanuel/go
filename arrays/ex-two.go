package main

import "fmt"

func main() {
	nums := [...]int{30, -1, -6, 90, -6, -3}

	var count uint8

	for _, i := range nums {
		if i > 0 && i%2 == 0 {
			count++
		}
	}
	fmt.Println("Positive even count:", count)
}
