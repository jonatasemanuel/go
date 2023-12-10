package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := string(99) // output: "c" | in ascii
	fmt.Println(s)

	// s1 := string(44.2) cannot convert a float to string

	var myStr = fmt.Sprintf("%.3f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 12)
	fmt.Println(myStr1)

	fmt.Println(string(234))

	/* STRING TO NUMBERS */
	s1 := "3.123"
	fmt.Printf("%T\n", s1) // string

	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Printf("%T %f\n", f1, f1)

	i, err := strconv.Atoi("-50") // Ascii-to-int
	s2 := strconv.Itoa(20)        // Int-to-ascii

	fmt.Printf("i type is %T, i value is %v\n", i, i)
	fmt.Printf("s2 type is %T, s2 value is %v\n", s2, s2)
}
