/*  and launch 50 goroutines that calculate concurrently the square root of all the numbers between 100 and 149 (both included). */

package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(50)
	defer wg.Wait()
	for i := 100.; i <= 149; i++ {
		go func(n float64, wg *sync.WaitGroup) {
			x := math.Sqrt(n)
			fmt.Println(x)

			wg.Done()
		}(i, &wg)
	}
}
