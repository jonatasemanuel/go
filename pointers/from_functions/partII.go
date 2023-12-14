package main

import "fmt"

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 8
	price = 500.4
	name = "Mobile Phone"
	sold = false

}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 4
	*price = 420.4
	*name = "PikaPhone"
	*sold = false

}

type Product struct {
	price       float64
	productName string
}

func changeProduct(p Product) {
	p.price = 300
	p.productName = "BMX"
}

func changeProductByPointer(p *Product) {
	(*p).price = 300 // equivalent to p.price = 300
	p.productName = "BMX"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

// A map value is already a pointer **
func changeMap(m map[string]int) {
	m["a"] = 10
	m["c"] = 20
	m["b"] = 30

}

func main() {
	quantity, price, name, sold := 5, 234.23, "Laptop", true
	fmt.Println("Before:", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("After:", quantity, price, name, sold)

	// TO CHANGE A VALUE WITH FUNC WE NEDD PASS THE ADDRESS(pointer) AS func() ARGUMENTS
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("After By Pointers:", quantity, price, name, sold)

	gift := Product{
		price:       100,
		productName: "Watch",
	}

	changeProduct(gift)
	fmt.Println("Before:", gift)

	changeProductByPointer(&gift)
	fmt.Println("After:", gift)

	//** SLICES AND MAPS (don't need to use a pointer)**//
	/* They don't store the actual data,
	but a reference to another memory address where the data is stored
	*/

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println("Prices slice after calling changeSlice():", prices)

	myMap := map[string]int{"a": 1, "b": 2}
	changeMap(myMap)
	fmt.Println("my Map after calling changeMap():", myMap)

	// ARRAYS BEHAVE JUST LIKE INTS FLOATS OR STRINGS
	// BTW, it's not good practice to pass arrays to functions, pass slice instead
	// Pass pointers to functions only when it's really necessary
}
