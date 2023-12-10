package main

import "fmt"

type mile float64
type kilometer float64

const m2km = 1.609 // 1mile==1.608km

func main() {
	var mileBerlimToParis mile = 655.3
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlimToParis * m2km)
	fmt.Printf("%.3f km", kmBerlinToParis)
}
