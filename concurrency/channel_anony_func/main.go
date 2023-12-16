package main

import (
	"fmt"
	"net/http"
	"runtime"

	"time"
)

func checkUrl(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		// fmt.Println(err)

		s := fmt.Sprintf("%s is DOWN!!!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		fmt.Println(s)
		c <- url // Sending into the channel
	} else {

		s := fmt.Sprintf("%s -> Status Code: %d \n", url, resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)
		fmt.Println(s)
		c <- url
	}

}

func main() {

	urls := []string{"https://golang.org", "https://google.com", "https://medium.com"}

	c := make(chan string)
	for _, url := range urls {
		go checkUrl(url, c)
	}

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	/* 	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	} */

	/* for {
		go checkUrl(<-c, c) // func return url trowutfhf the channel
		fmt.Println(strings.Repeat("_-", 30))
		time.Sleep(time.Second * 2)
	} */

	/* for url := range c {
		time.Sleep(time.Second * 2)
		go checkUrl(url, c)
	} */

	for url := range c {
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkUrl(u, c)
		}(url) // a url copy

	}
}
