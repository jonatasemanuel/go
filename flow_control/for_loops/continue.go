package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue // back to condition
		}
		fmt.Println(i)
	}
}
