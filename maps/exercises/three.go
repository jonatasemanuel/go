/*
	Consider the following map declaration: m := map[int]bool{"1": true, 2: false, 3: false}

1. The above map declaration has an error. Solve the error!
2. Delete a key:value pair from the map.
3. Iterate over the map and print out each key and value.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	m := map[int]bool{1: true, 2: false, 3: false}
	cP := make(map[int]bool)
	delete(m, 2)

	for i, v := range m {
		fmt.Printf("Key: %d Value: %v\n", i, v)
		cP[i] = v
	}

	fmt.Println(strings.Repeat("-=", 10))

	cP[1] = false
	cP[23] = true
	for i, v := range cP {
		fmt.Printf("Key: %d Value: %v\n", i, v)
	}

	fmt.Println(strings.Repeat("-=", 10))

	for i, v := range m {
		fmt.Printf("Key: %d Value: %v\n", i, v)
		cP[i] = v
	}
}
