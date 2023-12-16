package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i=", i)
	}
	fmt.Println("f1 execution finished")

}
func f2() {
	fmt.Println("f2 (goroutine) execution started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 execution finished")

}
func main() {
	fmt.Println("Main execution started")
	fmt.Println("No. of CPU's:", runtime.NumCPU())
	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	// GOOS and GOARCH are environment variables
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("Arch:", runtime.GOARCH)

	//  GOMAXPROCS sets the maximum number of CPUs that can be executing simultaneously and returns
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	go f1()
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine())

	go f2()

	time.Sleep(time.Second * 2)
	fmt.Println("Main execution stopped")
}
