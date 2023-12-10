package main

import "fmt"

// func main() {
// 	primes := [6]int{2, 3, 5, 7, 11, 13}
//
// 	var s []int = primes[1:4]
// 	fmt.Println(s)
// }

// SLICES POINTERS
// func main() {
// 	names := [4]string{
// 		"John",
// 		"Paul",
// 		"George",
// 		"Ringo",
// 	}
// 	fmt.Println(names)
//
// 	a := names[0:2] // john paul
// 	b := names[1:3] // paul george
// 	fmt.Println(a, b)
//
// 	b[0] = "XXX"
// 	fmt.Println(a, b) // john xxx -> xxx george
// 	fmt.Println(names)
// }

// SLICE LITERALS

//	func main() {
//		q := []int{2, 3, 5, 7, 11, 13}
//		fmt.Println(q)
//
//		r := []bool{true, false, true, true, false, true}
//		fmt.Println(r)
//
//		s := []struct {
//			i int
//			b bool
//		}{
//			{2, true},
//			{3, false},
//			{5, true},
//			{7, true},
//			{11, false},
//			{13, true},
//		}
//		fmt.Println(s)
//	}
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
