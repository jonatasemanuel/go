package main

import "fmt"

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {
	var x interface{}
	x = cube{edge: 5}
	ball, ok := x.(cube)
	if ok {
		fmt.Printf("Cube Volume: %v\n", volume(ball))
	}

}
