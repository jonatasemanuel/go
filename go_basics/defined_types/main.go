package main

import "fmt"

/*
WHY DEFINE NEW TYPES ?
  - Attach methods to newly difined types
  - Type safety: must convert one type into another to perform operations with them
  - Readability: when we defined a new type let's say *type usd float64* we know that new
    type represents the US Dollar not only floats.
*/
type km float64
type mile float64

func main() {
	type age int        // int its underlying type
	type oldAge age     // int its underlying type
	type veryOldAge age // int its underlying type

	type speed uint
	var s1 speed = 10
	var s2 speed = 20
	fmt.Printf("%v\n", s2-s1) // same type

	var x uint
	// x = s1 -> error, different types
	x = uint(s1) // now works
	_ = x

	var s3 speed = speed(x)
	_ = s3

	var parisToLondon km = 465
	var distanceInMiles mile

	distanceInMiles = mile(parisToLondon) / 0.621 // convert into left type
	fmt.Printf("%.2f", distanceInMiles)
}
