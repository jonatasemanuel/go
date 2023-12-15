/*
 1. Create a function called sum() that calculates and then prints out  the sum of 2 float numbers it receives as arguments.

Format the result with 2 decimal points.
2. From main launch 3 goroutines that execute the function you have just created (sum)
3. Synchronize the goroutines and the main function using WaitGroups
*/
package main

import (
	"fmt"
	"sync"
)

func sum(a, b float64, gw *sync.WaitGroup) {
	defer gw.Done()
	rs := a + b
	fmt.Printf("%.2f \n", rs)
}

func main() {
	var gw sync.WaitGroup
	gw.Add(3)
	defer gw.Wait()
	go sum(3.2, 1.8, &gw)
	go sum(.2, 1.3, &gw)
	go sum(2.2, 5.1, &gw)
}
