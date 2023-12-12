package main

import "fmt"

func main() {
	friends := map[string]int{"Jordan": 23, "Messi": 30, "CR": 7, "R": 9}
	legends := make(map[string]int)

	for k, v := range friends {
		legends[k] = v
	}

	legends["Eren"] = 1
	fmt.Printf("%#v ---| %#v\n", legends, friends)
}
