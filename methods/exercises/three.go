/*
 1. Create a new struct type called book with 2 fields: price and title of type float64 and string.

2. Create a method for the newly defined type called vat that returns the vat value if vat is 9%.
3. Create a method called discount that takes a book value as receiver and decreases the price of the book by 10%.
*/
package main

import "fmt"

type book struct {
	price float64
	title string
}

func (v *book) vat() float64 {
	return v.price * 0.9
}

func (d *book) discount() {
	d.price = d.price * 0.9
}

func main() {
	jornal := book{price: 30., title: "A linguagem C"}
	fmt.Printf("%#v\n", jornal)
	jornal.discount()
	fmt.Printf("%#v\n", jornal)
	fmt.Println(jornal.vat())
}
