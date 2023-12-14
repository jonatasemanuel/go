package main

import "fmt"

type booky struct {
	title string
	price float64
}

// method for book type
func (b *booky) changePrice(p float64) {
	b.price = p
}

func main() {
	// book value
	bestBook := booky{title: "The Trial by Kafka", price: 9.9}

	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price) // no change
}
