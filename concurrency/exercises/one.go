package main

import (
	"fmt"
	"sync"
)

func sayHello(n string, gr *sync.WaitGroup) {
	defer gr.Done()
	fmt.Printf("Hello, %s!\n", n)
}

func main() {
	var gr sync.WaitGroup
	defer gr.Wait()
	gr.Add(1)

	go sayHello("Mr. Wick", &gr)
}

/* Your task is to synchronize main and the goroutine using WaitGroups. The program should print the string received as argument by sayHello(). */
