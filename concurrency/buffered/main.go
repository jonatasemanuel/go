package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after\n", i)
		}
		close(c)
		// BTW, doesn't need to close every channel
		// only necessary when it is important to tell the receiving goroutines that all have been sent.
	}(c)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c { // v := <- c
		fmt.Println("main goroutine received value from channel", v)
	}

	// UNBUFFERED GIVE STRONGER SYNC
	// WITH BUFFERED CHANNELS, THESE OPERATIONS ARE DECOUPLED | ALSO WHEN WE KNOW AN UPPER BOUND  ON THE NUMBER OF VALUES WILL BE SENT ON A CHANNEL
}
