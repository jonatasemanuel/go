package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}
	fmt.Printf("x:%d y:%.1f z:%s score:%#v\n", x, y, z, score)
	fmt.Printf("%q\n", z)
	fmt.Printf("x:%v y:%v z:%v\n", x, y, z)
	fmt.Printf("y:%T score:%T\n", y, score)
}
