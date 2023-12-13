/*
 1. Create your own struct type called person that stores the following data: name, age and a list with favorite colors.

2. Declare and initialize two values of type person, one called me and another called you.
3. Print out the struct values.
*/
package main

import "fmt"

func main() {
	type Grades struct {
		grade  int
		course string
	}

	type Person struct {
		name           string
		age            int
		favoriteColors []string
		myGrades       Grades
	}

	me := Person{
		name:           "Joe",
		age:            24,
		favoriteColors: []string{"Red", "Yellow"},
		myGrades:       Grades{98, "Golang"},
	}
	u := Person{
		name:           "Yu Yu",
		age:            14,
		favoriteColors: []string{"Green", "Blue"},
	}

	fmt.Println(me, u)
	/*
	   	Consider the code from the previous exercise and:

	   1. Change the name or the struct value called me to "Andrei"
	   2. Take in a new variable the favorites colors of struct value called you. Print out the type and the value of the new variable.
	   3. Add a new favorite color to the second struct value called you.
	   4. Print out the struct values.
	*/
	me.name = "Andrei"

	var colorU []string = u.favoriteColors
	fmt.Printf("type: %T Value: %v\n", colorU, colorU)
	colorU = append(colorU, "Purple")
	u.favoriteColors = colorU
	fmt.Println(u)

	/* Iterate and print out the favorite colors of the struct value called me. */

	for _, color := range me.favoriteColors {
		fmt.Println(color)
	}

	/* 1. Create a new struct type called grades with  2 fields: grade and course
	2. Add another field of type grades to person struct type (embedded struct).
	3. Change the initialization of the struct values called me and you to contain information about the grades.
	4. Change the grade and the course of one struct value to "Golang" and 98.
	5. Print out the struct values. */

	fmt.Println(me)
}
