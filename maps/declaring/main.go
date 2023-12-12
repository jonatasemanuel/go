package main

import "fmt"

func main() {
	var employees map[string]string // nil
	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of pairs: %d\n", len(employees))

	fmt.Printf("The value for key %q is %q", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	var m1 map[[4]int]string // array as index
	_ = m1

	// employees["Dan"] = "Programmer" -> err: not initialize

	people := map[string]float64{}

	people["John"] = 23.1
	people["Marry"] = 4.20
	fmt.Println(people)

	mapOne := make(map[int]int)
	mapOne[4] = 8

	balances := map[string]float64{
		"USD": 323.12,
		"BRL": 420.89,
		"CHF": 3234.1,
	}
	fmt.Println(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	balances["USD"] = 2145.
	balances["EUR"] = 123.2
	fmt.Println(balances)

	fmt.Println(balances["RON"])

	v, ok := balances["RON"]
	if ok {
		fmt.Println("The RON balance is: ", v)
	} else {
		fmt.Println("The RON Key doesn't exist in the map!")
	}

	for k, v := range balances {
		fmt.Printf("Key: %#v - Value: %#v\n", k, v)
	}
	delete(balances, "BRL")
	fmt.Println(balances)

	// MAP CAN ONLY BE COMPARE WITH NIL

	a := map[string]string{"A": "X", "B": "Cu"}
	b := map[string]string{"A": "X", "B": "Cu"}

	// TURNING MAPS INTO A STRING, NOW WE CAN COMPARE
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	fmt.Println(s1, s2)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

}
