package main

import "fmt"

func main() {
	const (
		sunToEarth = 149_600_000 * 1000 // (1km == 1000m)
		speedOf    = 299_792_458
	)

	const howLong = sunToEarth / speedOf
	fmt.Printf("Along of %v seconds", howLong)
}
