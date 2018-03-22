package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"time"
)

type pingStatus struct {
	url    string
	status string
}

func main() {
	ch := make(chan *pingStatus)
	urls, _ := readLines("urls.txt")

	for _, url := range urls {
		go ping(url, ch)
	}

	for {
		pingStatus := <-ch
		go func(url string) {
			time.Sleep(5 * time.Second)
			ping(url, ch)
		}(pingStatus.url)

		log.Printf("%v - %v", pingStatus.url, pingStatus.status)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	var urls []string

	for fs.Scan() {
		urls = append(urls, fs.Text())
	}

	return urls, nil
}

func ping(url string, ch chan *pingStatus) {
	ps := &pingStatus{url: url}

	resp, err := http.Get(url)
	if err != nil {
		ps.status = err.Error()
		ch <- ps
		return
	}

	ps.status = resp.Status
	ch <- ps
}
