package main

import "fmt"

func main() {
	var m1 map[int]bool
	// m1[5] = true
	_ = m1

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	fmt.Println(m2 == nil, m3)
}
