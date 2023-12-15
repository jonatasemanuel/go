/*
Create a function literal (a.k.a. anonymous function)
that sends the string value if receives as argument to main func using a channel.
*/
package main

import "fmt"

func main() {

	ch := make(chan string)
	go func(s string) {
		ch <- s
	}("Test")

	rs := <-ch
	fmt.Printf("%T %v\n", rs, rs)
}
