/*
 1. Declare a map called m1 which value is nil. Print out its type and value.

2. Declare a map called m2. It's keys are of type int and values of type string.  Initialize the map using  a map literal with two key:value pairs.
3. Add the following key: value to the map: 10: "Abba"
4. Retrieve the value of an existing key and the value of a non existing key. Think about the results.
*/
package main

import "fmt"

func main() {
	var m1 map[string]float64
	fmt.Printf("%T: %#v\n", m1, m1)

	m2 := map[int]string{777: "Raff", 23: "Jordan"}
	m2[10] = "Abba"
	fmt.Printf("%#v \n", m2[2414]) // ""
}
