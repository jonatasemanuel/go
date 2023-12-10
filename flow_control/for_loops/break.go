package main

import "fmt"

func main() {
	count := 0

	for i := 0; true; i++ { // true to infinite loop
		if i%13 == 0 {
			fmt.Printf("%d : %d is divisible by 13 \n", count+1, i)
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("Done!")
}
