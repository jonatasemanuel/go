package main

import "fmt"

func main() {
	s1 := "Hi there Go!" // string literal
	fmt.Println(s1)

	fmt.Println("He seys: \"hello\"")

	fmt.Println(`He seys: "Hello!"`) // raw string

	s2 := `I like \n Go!` // raw string -> printing everything
	fmt.Println(s2)

	fmt.Println("Price: 100\nBrand: Nike")
	fmt.Println(`
Price: 100
Brand: Nike
	`)

	fmt.Println(`C:\Users\Joe`)
	fmt.Println("C:\\Users\\Joe") // scape backslach

	var s3 = "I love " + "Go" + "Programing"
	fmt.Println(s3 + "!")

	fmt.Println("Element at index 0:", s3[0])

	// s3[5] = 'x' -> strings are imutable
	fmt.Printf("%s\n", s3)
	fmt.Printf("%q\n", s3)
}
