/*
 1. Using single function creates a file called info.txt in the current directory. If the file already exists it will truncate it to zero size.

2. Write the string "The Go gopher is an iconic mascot!" to the file.
*/
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"info.txt",
		os.O_TRUNC|os.O_WRONLY|os.O_CREATE,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferWrited := bufio.NewWriter(file)
	byteSlice := []byte("The Go gopher is an iconic mascot!")
	byteWritten, err := bufferWrited.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %#v\n", byteWritten)

	bufferWrited.Flush()
}
