package main

import "fmt"

func main() {
	const (
		Jun = (iota + 6)
		Jul
		Aug
	)

	fmt.Printf("%v", Jun, Jul, Aug)
}
