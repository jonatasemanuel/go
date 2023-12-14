/*
	Consider the following type and interface declaration.

1. Create a Go program where car type implements the vehicle interface.
2. Create a variable of type vehicle that holds a car struct value.
3. Call the methods (Licence and Name) on the interface value declared at step 2
4. Run the program without errors.
*/
package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}
func (c car) Name() string {
	return c.brand
}

func main() {
	var gol vehicle = car{licenceNo: "Yes", brand: "WM"}
	fmt.Println(gol.Name())
	fmt.Println(gol.License())
}
