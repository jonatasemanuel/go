/*
Using a composite literal declare and initialize a slice of type string with 3 elements.
Iterate over the slice and print each element in the slice and its index.
*/
package main

import "fmt"

func main() {
	names := []string{"Joe", "Jowzee", "Joey"}
	for i, name := range names {
		fmt.Printf("%d: %s\n", i, name)
	}

}
