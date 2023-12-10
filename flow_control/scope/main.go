package main

import "fmt"
import f "fmt" // permitted in go by alias

const done = false // package scoped

var b int = 10 // can be used also in another file in the same package

func main() {
	var task = "Running" // local
	fmt.Println(task, done)

	const done = true // local
	f.Printf("Done in main() is %v\n", done)
	f1()

	f.Println("Bye bye!")
}

func f1() {
	fmt.Printf("Done in f1(): %v\n", done) // package scope
}
