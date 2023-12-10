package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"
	switch language {
	case "Python":
		fmt.Println("You're learning Python! You don't use curly braces but indentations!!")
		// don't need use break
	case "Go", "golang":
		fmt.Println("Good, go for Go! You are using curly braces {}!")

	default:
		fmt.Println("Any other suck language is a good start!")
	}

	n := 5
	switch true { // our just switch
	case n%2 == 0:
		fmt.Println("Even number!")
	case n%2 != 0:
		fmt.Println("Odd number!")
	default:
		fmt.Println("Never here!")
	}

	hour := time.Now().Hour()
	// fmt.Println(hour)
	switch {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}
}
