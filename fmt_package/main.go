package main

import "fmt"

func main() {
	// "fmt.Printf" -> to formatted output
	// %d :int | %f :float  %.2f  | %s :string | %q :quoted strigns **d comes from decimal
	// %v -> replaced by any type
	// %T (uppercase) -> type of variable
	// %t -> bool
	// %b -> base 2 | binary
	// %x -> hexadecimal
	name := "Joe"
	fmt.Println("Hello, my name is ", name)

	a, b := 4, 6
	fmt.Println("Sum: ", a+b, "Mean Value: ", (a+b)/2)

	fmt.Printf("Your age is %d\n", 21)

	fmt.Printf("x is %d, y is %f\n", 5, 6.8)

	// Double quotes
	fmt.Printf("He says: \"Hello Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.141559
	_, _ = figure, pi

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)

	fmt.Printf("Pi constant is %f\n", pi)

	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n", figure, radius, float64(radius)*2*pi) // both type needs to be the same to do an operation

	fmt.Printf("This is a %q", figure)

	fmt.Printf("figure is of type %T \n", figure)
	fmt.Printf("radius is of type %T \n", radius)
	fmt.Printf("pi is of type %T \n", pi)

	closed := true
	fmt.Printf("File closed: %t\n", closed)

	fmt.Printf("%b \n", 55)
	fmt.Printf("%08b \n", 55) // To display 8bits -> show left zeros

	fmt.Printf("%x \n", 100)

	x := 3.4
	y := 6.9

	fmt.Printf("x * y = %.3f\n", x*y)
}
