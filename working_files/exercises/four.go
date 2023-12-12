package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)

}
