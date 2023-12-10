package main

import "fmt"

// type Vertex struct {
// 	X int
// 	y int
// }
//
// func main() {
// 	v := Vertex{1, 3}
//
// 	// Struct fields are accessed using a dot.
// 	v.X = 4
// 	fmt.Println(v.X)
// }

// POINTERS TO STRUCTS

// type Vertex struct {
// 	X int
// 	Y int
// }

// func main() {
// 	v := Vertex{1, 2}
// 	p := &v
// 	p.X = 1e9
// 	fmt.Println(v)
// }

// STRUC LITERALS

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
