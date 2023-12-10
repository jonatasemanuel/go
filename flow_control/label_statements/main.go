package main

import "fmt"

/*
	- Label are used in "break", "continues" and "goto" statements
	- It's illegal to define a label that is never used
	- Labels arte not block scoped and do conflict with identifiers
	- Function body scope
	- Most of the time labels are used to terminate outer enclosing loops
*/

func main() {
	people := [5]string{"Helen", "Mark", "Brenda", "Antonio", "Micheal"}
	friends := [2]string{"Mark", "Marry"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next step...")

}
