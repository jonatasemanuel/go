package main

import "fmt"

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice

}

/*
-	We cannot pass a pointer type in method receiver
type distance *int

func (d *distance) m1() {
	fmt.Println("tete")

}

*/

func main() {
	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "Fusca", 23000)
	fmt.Println(myCar)

	myCar.changeCar1("Gol", 23111)
	fmt.Println(myCar)

	// pointer receiver
	(&myCar).changeCar2("Golf", 45666)
	fmt.Println(myCar)
	myCar.changeCar2("Vectra", 23433) // & implicit
	fmt.Println(myCar)

	// Valid ways
	var yourCar *car
	yourCar = &myCar

	fmt.Println(*yourCar)

	yourCar.changeCar2("VW", 34000)
	fmt.Println(*yourCar)
}
