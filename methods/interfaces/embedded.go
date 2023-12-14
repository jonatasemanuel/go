package main

import (
	"fmt"
	"math"
)

/* type shapi interface {
	area() float64
}

type object interface {
	volume() float64
}
*/

// Geomtry is embedding shape and object interfaces
type geometry interface {
	area() float64
	volume() float64
	getColor() string
}

type cube struct {
	edge  float64
	color string
}

func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)

}

func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

func (c cube) getColor() string {
	return c.color
}

func measure(g geometry) (float64, float64, string) {
	a := g.area()
	v := g.volume()
	gC := g.getColor()
	return a, v, gC

}

func main() {

	c := cube{edge: 2, color: "Red"}
	a, v, cG := measure(c)
	fmt.Printf("Area: %v | Vol: %v³ | Color: %s\n", a, v, cG)
}
