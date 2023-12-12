package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file := "learn.txt"
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist!")
		}
	}

	err = os.Remove(file)
	if err != nil {
		log.Fatal(err)
	}

}
