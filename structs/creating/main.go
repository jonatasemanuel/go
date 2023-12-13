/*
  - A STRUCT is a sequence of named elemtns, called fields. Each of them has a name and a type.
  - The struct fields are like the instance attributes we define in OOP
  - Go does NOT have a class-object architecture.
  - A structs is a schema containing a blueprint of data a structure will hold.
    This blueprint is fixed at compile time. It's not allowed to change the name or the
    type of the fields at runtime. We can't add or remove fields from struct at runtime.
*/
package main

import "fmt"

func main() {

	title1, author1, years1 := "The Devine Comedy", "Dante Aligheri", 1320
	title2, author2, years2 := "Macbeth", "Willian Shakespeare", 1606

	type book struct {
		title, author string
		year          int
	}

	bookOne := book{}
	bookOne.title = title1
	bookOne.author = author1
	bookOne.year = years1

	bookTwo := book{title2, author2, years2} // Struct order
	// best call
	bookThree := book{
		title:  "Animal Farm",
		author: "George Orwell",
		year:   1945,
	}

	fmt.Println(bookOne, bookTwo, bookThree)

	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)

	lastBook.author = "Leo Tolstoy"
	lastBook.year = 1878
	fmt.Printf("%+v\n", lastBook)

	aBook := book{title: "Anna Karenina", author: "Leo Tolstoy", year: 1878}
	fmt.Println(aBook == lastBook) // true, have some fields

	// CLONE

	newBook := aBook
	newBook.year = 2020
	fmt.Println(newBook, aBook)

	// ANONYMOUS STRUCTS
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}
	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana's Age: %d\n", diana.age)

	// ANONYMOUS FIELDS
	type Book struct {
		string
		float64
		bool
	}
	b1 := Book{"1984 by G. Orwell", 10.2, false}
	fmt.Printf("%#v\n", b1) // the types turn into fields name.
	fmt.Printf(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 2345, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e)
}
