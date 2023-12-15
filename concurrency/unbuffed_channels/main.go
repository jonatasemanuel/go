package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) // unbuffered channel

	// If the capacity is non-zero, make creates a buffed channel
	c2 := make(chan int, 3)
	_ = c2

	go func(c chan int) {
		fmt.Println("func goroutines starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutines after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	time.Sleep(time.Second * 2)
}
