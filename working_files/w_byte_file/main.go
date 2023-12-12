package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlice := []byte("I learn Golang!")
	byteWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", byteWritten)

	// os.WriteFile() -> handle creatign, opening, writing a slice of bytes and close the f
	bs := []byte("Go Programming is cool!")
	err = os.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", byteWritten)

}
