package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Stat("info-x.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist!")
		}
	}

	oldFile := "info-x.txt"
	newFile := "learn.txt"
	err = os.Rename(oldFile, newFile)
	if err != nil {
		log.Fatal(err)
	}
}
