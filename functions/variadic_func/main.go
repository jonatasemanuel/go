package main

import (
	"fmt"
	"strings"
)

func fOne(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func fTwo(a ...int) {
	a[0] = 50
}

func SumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.
	for _, v := range a {
		sum += v
		product *= v
	}
	return sum, product
}

func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")

	reutrnString := fmt.Sprintf("Age %d, Full name: %s", age, fullName)
	return reutrnString

}
func main() {
	fOne(1, 2, 3, 4, 5)
	fOne()

	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)

	// Pass a slice to a variadic function by post fixing it with the Variadic Operator
	fmt.Println(nums)
	fOne(nums...)

	// Part II
	fTwo(nums...)
	fmt.Println("Nums:", nums)

	s, p := SumAndProduct(2., 5., 12., 8., 45.2)
	fmt.Println(s, p)

	info := personInformation(30, "Wolfgang", "Amadeus", "Mozart")
	fmt.Println(info)
}
