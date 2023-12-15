package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func funcOne(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second)
	}
	fmt.Println("f1 execution finished")
	wg.Done()
}
func funcTwo() {
	fmt.Println("f2 (goroutine) execution started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution finished")

}
func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go funcOne(&wg)
	fmt.Println("No. of Goroutines after go funcOne():", runtime.NumGoroutine())

	go funcTwo()

	wg.Wait()
	fmt.Println("Main execution stopped")
}
