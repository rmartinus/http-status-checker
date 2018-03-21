package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
)

type pingStatus struct {
	url    string
	status string
}

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	ch := make(chan *pingStatus)
	n := 0
	for fs.Scan() {
		go ping(fs.Text(), ch)
		n = n + 1
	}

	for i := 0; i < n; i++ {
		pingStatus := <-ch
		log.Printf("%v - %v", pingStatus.url, pingStatus.status)
	}
	log.Println("Done")
}

func ping(url string, ch chan *pingStatus) {
	ps := &pingStatus{url: url}
	log.Printf("Pinging %v", url)

	resp, err := http.Get(url)
	if err != nil {
		ps.status = err.Error()
		ch <- ps
		return
	}

	ps.status = resp.Status
	ch <- ps
}
