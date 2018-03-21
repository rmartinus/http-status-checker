package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	ch := make(chan string)
	n := 0
	for fs.Scan() {
		go ping(fs.Text(), ch)
		n = n + 1
	}

	for i := 0; i < n; i++ {
		log.Print(<-ch)
	}
	log.Println("Done")
}

func ping(url string, ch chan string) {
	log.Printf("Pinging %v", url)

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%v - ERROR: %v\n", url, err)
		return
	}

	ch <- fmt.Sprintf("%v - %v\n", url, resp.Status)
}
