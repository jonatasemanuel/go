package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/* fmt.Println("os.Args", os.Args)
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st arg: ", os.Args[1])
	fmt.Println("2st arg: ", os.Args[2])
	fmt.Println("No if inside os.Args: ", len(os.Args)) */

	var result, err = strconv.ParseFloat(os.Args[1], 64)
	_ = err
	fmt.Println(result)
}
