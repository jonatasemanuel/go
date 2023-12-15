/*
  - A channel in Go provides a connection between two go routines,
    allowing them to communicate.
  - Channels are used to communicate in between running go routines.
  - A channel is like a pointer, and passing channels to functions has the same effect.
  - A channel has to operations, Send and Receiver
  - Channel Operator <-
*/
package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch)

	ch = make(chan int)
	fmt.Println(ch)

	// Short declaration
	c := make(chan int)

	// Send
	c <- 10

	// Receive
	num := <-c

	fmt.Println(<-c)

	_ = num

	close(c) // close a channerl

}
