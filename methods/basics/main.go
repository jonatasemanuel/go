/*
	GO doesn't have class, but we can define methods on predifined types

- A type may have a method set associated with it which enhances the
type with extra behavior

- Methods belongs to a type, and function belongs to a package
*/
package main

import (
	"fmt"
	"time"
)

type names []string

// n is a method receiver
func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}

}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)
	fmt.Printf("%v\n", day)

	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds)
	fmt.Printf("Seconds in a day: %v\n", seconds)

	friends := names{"Marry", "Jane"}
	friends.print()

	names.print(friends)

	var n int64 = 98348533
	fmt.Println(n)
	fmt.Println(time.Duration(n))
}
