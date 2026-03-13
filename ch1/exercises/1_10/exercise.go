package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// Find a website that produces a large a mount of data. Investigate caching
// by running fetch all twice in succession to seee whether reported time changes much.
// Do you get the same content each time? Modify fetchall to print
// its output to a file so that it can be examined

// To Save the Output, Make sure you pass you Path to Project in the place holder variable {PATH_TO_PROJECT}

// In this project there are some weaknesses
// 1. When one go route fails the entire program panics.
func main() {
	startTime := time.Now()
	urls := os.Args[1:]
	if len(urls) == 0 {
		log.Fatal("Please provide valid URLs")
	}
	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("Time elapsed : %.2fs", time.Since(startTime).Seconds())

}

func fetch(url string, ch chan string) {
	var finalUrlString string = url
	if !strings.HasPrefix("https://", url) && !strings.HasPrefix("http://", url) {
		finalUrlString = "http://" + finalUrlString
	}
	startTime := time.Now()
	response, err := http.Get(finalUrlString)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching url %s. Err : %v", url, err)
	}

	filePath := fmt.Sprintf("{PATH_TO_PROJECT}/ch1/exercises/1_10/output%d", time.Now().Nanosecond())

	file, fileErr := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if fileErr != nil {
		ch <- fmt.Sprintf("Error opening file err: %v", fileErr)
	}

	byteCount, copyErr := io.Copy(file, response.Body)
	response.Body.Close()

	if copyErr != nil {
		ch <- fmt.Sprintf("Errror in copying output to file %v", copyErr)
	}

	ch <- fmt.Sprintf("Url: %s\t Time Elapse: %.2f Response Size: %d", url, time.Since(startTime).Seconds(), byteCount)
}
