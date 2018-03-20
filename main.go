package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{"https://golang.org/", "https://www.google.com.au/"}

	for _, url := range urls {
		go checkStatus(url)
	}

	fmt.Scanln()
	fmt.Println("done")
}

func checkStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(url, resp.Status)
}
