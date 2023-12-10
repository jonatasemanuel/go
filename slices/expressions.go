package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	// a[start:stop]   | [:3] 0 -> 3 [1:] 1 -> max len
	b := a[1:3] // [2 3]
	fmt.Printf("%v, %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}

	s2 := s1[1:4]
	fmt.Println(s2)

	fmt.Println(s1[2:]) // [3, 4 , 5, 6] : same as s1[2:len(s1)]
	fmt.Println(s1[:3]) // same as [0:3]
	fmt.Println(s1[:])

	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	s1 = append(s1[:4], 200)
	fmt.Println(s1)
}
