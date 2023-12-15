package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// CLI FLAG TO DETECTOR: go run -race [main.go]
	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	// 1. Mutex value
	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			/* 2. The are two methods type lock and unlock
			- Any code that is present between a call to lock and unlock
			will be executed by only one routine */
			m.Lock()

			n++

			m.Unlock()
			wg.Done()
		}()
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()

			n--
			wg.Done()
		}()

	}
	wg.Wait()

	fmt.Println("Final value of n:", n)
}
