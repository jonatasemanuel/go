package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error

	newFile, err = os.Create("a.txt")
	if err != nil {
		// fmt.Println(err)
		// os.Exit(1)
		log.Fatal(err)
		// log is thread safe, also adds timing information automatically.
	}

	// change de size of file name "a.txt"
	err = os.Truncate("a.txt", 0)
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()

	// OPENING A FILE WITH MORE OPTIONS
	/*
		We can Use opening attributes individually or combined
		using an OR between them
		e.g. os.O_CREATE|os.O_APPEND
		or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
		os.O_RDONLY // Read only
		os.O_WRONLY // Write only
		os.O_RDWR // Read and write
		os.O_APPEND // Append to end of file
		os.O_CREATE // Create is none exist
		os.O_TRUNC // Truncate file when opening
	*/

	// file, err := os.Open("a.txt")
	// file, err := os.OpenFile("a.txt", os.O_APPEND, 0644)
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0655) // create ou append
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// FILE INFO
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")

	p := fmt.Println

	p("File:", fileInfo.Name())
	p("Size [bytes]:", fileInfo.Size())
	p("Last modified:", fileInfo.ModTime())
	p("Is Directory?", fileInfo.IsDir())
	p("Permissions:", fileInfo.Mode())

	// CHECKING IF FILE EXIST
	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist!")

		}
	}

	// RENAME FILE
	/* oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
	file.Close() */

	err = os.Remove("a.txt")
	if err != nil {
		log.Fatal(err)
	}
}
