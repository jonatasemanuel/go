package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Text() for string or Bytes() for byte slice
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords) // word by word
	scanner.Split(bufio.ScanRunes) // word by word

	success := scanner.Scan() // bool type
	if success == false {
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it reached EOF")
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("First line found:", scanner.Bytes())
	fmt.Println("First line found:", scanner.Text())

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal()
	}

}
