package main

import "fmt"

// A DEFER STATEMENT DEFERS THE EXECUTION OF A FUNCTION UNTIL THE SURROUNDING FUNCTION RETURNS.
// func main() {
// 	defer fmt.Println("world")
//
// }
// 	fmt.Println("hello")

// Output: "hello" "world"

// STACKING DEFERS

// Deferred function calls are pushed onto a stack.
//When a function returns, its deferred calls are executed in last-in-first-out order.
//To learn more about defer statements read this blog post.

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
