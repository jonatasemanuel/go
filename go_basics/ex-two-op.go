package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	is := strconv.Itoa(i)
	fmt.Printf("%s %T\n", is, is)

	s2in, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("%d %T\n", s2in, s2in)
	} else {
		fmt.Println("Can not convert string to int")
	}

	f2 := fmt.Sprintf("%.2f", f)
	fmt.Printf("%s %T\n", f2, f2)

	sD, err := strconv.ParseFloat(s1, 64)
	if err == nil {
		fmt.Printf("%v %T\n", sD, sD)
	}

}
