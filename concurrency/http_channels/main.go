package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		// fmt.Println(err)

		s := fmt.Sprintf("%s is DOWN!!!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		c <- s // Sending into the channel
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d \n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodyBytes, err := io.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1] // | http:// www.google.com
			file += ".html"

			s += fmt.Sprintf("Writing response body to %s\n", file)

			err = os.WriteFile(file, bodyBytes, 0664)
			if err != nil {
				// log.Fatal(err)
				s += "Error writing file\n"
				c <- s
			}
			s += fmt.Sprintf("%s is UP\n", url)
			c <- s
		}
	}

}

func main() {

	urls := []string{"https://golang.org", "https://google.com", "https://medium.com"}

	c := make(chan string)
	for _, url := range urls {
		go checkAndSaveBody(url, c)
		fmt.Println(strings.Repeat("-=", 20))
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}
