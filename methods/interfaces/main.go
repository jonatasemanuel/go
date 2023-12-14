package main

import (
	"fmt"
	"math"
)

// The types that implement the interface must implement all methods (polymorphism)
type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) valume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

/*
	 func printCircle(c circle) {
		fmt.Println("Shape:", c)
		fmt.Println("Area:", c.area())
		fmt.Println("Perimeter:", c.perimeter())
	}

	func printRectangle(r rectangle) {
		fmt.Println("Shape:", r)
		fmt.Println("Area:", r.area())
		fmt.Println("Perimeter:", r.perimeter())
	}
*/

func printGeo(s shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v\n", s.area())
	fmt.Printf("Perimeter: %v\n", s.perimeter())
}

func main() {

	/* circleOne := circle{radius: 5.}
	recTangle := rectangle{width: 2., height: 3.}

	printGeo(circleOne)
	printGeo(recTangle)
	*/

	/* var s shape // polymorphism

	fmt.Printf("%T\n", s)

	ball := circle{radius: 2.5}
	s = ball
	printGeo(s)
	fmt.Printf("%T\n", s)

	room := rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("%T\n", s) */

	// TYPE ASSERTIONS AND TYPE SWITCHES
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	// s.volume() -> interface
	ball, ok := s.(circle)
	if ok == true {
		fmt.Printf("Ball volume: %v\n", ball.valume())

	}

	//s = rectangle{width: 2.3, height: 2.5}
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)

	}
}
