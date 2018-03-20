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
	for fs.Scan() {
		go checkStatus(fs.Text())
	}

	fmt.Scanln()
	fmt.Println("DONE")
}

func checkStatus(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("%v - ERROR: %v\n", url, err)
		return
	}

	log.Printf("%v - %v\n", url, resp.Status)
}
