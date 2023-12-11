package main

import (
	"fmt"
	"strings"
)

func main() {

	p := fmt.Println
	result := strings.Contains("I love Go Programming!", "lovex")
	p(result)

	result = strings.ContainsAny("success", "xys")
	p(result)

	result = strings.ContainsRune("golang", 'g')
	p(result)

	n := strings.Count("cheese", "e")
	p(n)

	n = strings.Count("Five", "") // 5 len + 1
	p(n)

	p(strings.ToLower("GO PYTHON JAVA"))
	p(strings.ToUpper("gO PyThoN JAVA"))

	p("go" == "go")
	p("GO" == "go")

	p(strings.EqualFold("GO", "go")) // true

	myStr := strings.Repeat("ab", 10)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", 2)
	p(myStr)

	myStr = strings.Replace("192.168.0.1", ".", ":", -1)
	p(myStr) // all ocurrence

	myStr = strings.ReplaceAll("192.168.0.1", ".", ":")
	p(myStr) // all ocurrence

	s := strings.Split("a, b, c", ",")
	fmt.Printf("%T\n", s)
	fmt.Printf("%#v\n", s)

	s = strings.Split("Go for Go!", "")
	fmt.Printf("%#v\n", s)

	s = []string{"I", "learn", "Golang"}
	myStr = strings.Join(s, "-/-")
	p(myStr)

	myStr = "Orange Green \nBlue Yellow"
	fields := strings.Fields(myStr)
	fmt.Printf("Fields: %T\n", fields)
	fmt.Printf("Fields: %#v\n", fields)

	s1 := strings.TrimSpace("\t  Goodbye Windows, Welcome Linux! \n") // remove black space at leader
	fmt.Printf("%q\n", s1)

	s2 := strings.Trim("...Hello Gophers!!!?", ".!?") // trim -> .!?
	fmt.Printf("%q\n", s2)
}
