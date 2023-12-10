/*
- SLICE HEADER contains 3 fields:
1. The address of the backing array (pointer).
2. The length of the slice. len() returns it.
3. The capacity of the slice. cap() returns it.
  - Runtime representation of a slice
  - A nil slice doesn't a have backing array.
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
		Part I

		s1 := []int{10, 20, 30, 40, 50}
		s3, s4 := s1[0:2], s1[1:3]

		s3[1] = 600
		fmt.Println(s1)
		fmt.Println(s4)

		arr1 := [4]int{10, 20, 30, 40}
		slice1, slice2 := arr1[:2], arr1[1:3]

		arr1[1] = 2

		fmt.Println(arr1, slice1, slice2)
	*/

	//	Part II - Coping slice into a new var
	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[:2]...)

	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	s1 := []int{10, 20, 30, 40, 50}
	newSlice := s1[:3]
	fmt.Println(len(newCars), cap(newSlice))

	newSlice = s1[2:5] // {30, 40, 50}
	fmt.Println(len(newCars), cap(newSlice))

	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p %p\n", &s1, &newSlice)

	newSlice[0] = 1000
	fmt.Println("s1: ", s1)

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("Array's size in bytes: %d\n", unsafe.Sizeof(a))
	fmt.Printf("Slice's size in bytes: %d\n", unsafe.Sizeof(s))
}
