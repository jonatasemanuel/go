package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// BYTE SLICE

	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := make([]byte, 2) // show a slice with 2 bytes [He]

	numberBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of bytes read: %d\n", numberBytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	//////////////////////////////////////////////////////////

	// READ ENTIRE FILE

	file, err = os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as string: %s\n", data)
	fmt.Printf("Number of bytes read: %d\n", len(data))

	// QUICK WAY

	data, err = os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read: %s\n", data)
}
