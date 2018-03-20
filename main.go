package main

import (
	"bufio"
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
	checkStatus(fs)

	log.Println("Done")
}

func checkStatus(fs *bufio.Scanner) {
	jobs := make(chan string)
	done := make(chan bool)

	go ping(done, jobs)

	for fs.Scan() {
		url := fs.Text()
		jobs <- url
	}

	close(jobs)
	<-done
}

func ping(done chan bool, jobs chan string) {
	for {
		url, more := <-jobs
		if more {
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("%v - ERROR: %v\n", url, err)
				continue
			}

			log.Printf("%v - %v\n", url, resp.Status)
		} else {
			done <- true
			return
		}
	}
}
