package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano() / 1_000_000

	//** The select statement lets a go routine wait on multiple communication operations.
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "Salut!"
	}()

	time.Sleep(time.Second * 1)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)

		default:
			fmt.Println("no activity")
		}
	}
	end := time.Now().UnixNano() / 1_000_000

	fmt.Println(end - start)
}
