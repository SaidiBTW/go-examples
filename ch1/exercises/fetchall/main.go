package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//Fetchall Fetches URLS in parallel and reports their time and sizes

func main() {
	urls := os.Args[1:]
	// Error check for invalid program invocation

	if len(urls) == 0 {
		log.Fatal("Please provide valid URLs")
	}
	// Channel for output
	ch := make(chan string)

	startTime := time.Now()

	for _, url := range urls {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2f elapsed\n", time.Since(startTime).Seconds())
}

func fetch(url string, outputChannel chan string) {
	startTime := time.Now()

	response, err := http.Get(url)
	if err != nil {
		outputChannel <- fmt.Sprintf("Error fetching url %s, %v", url, err)
	}

	nBytes, copyErr := io.Copy(io.Discard, response.Body)
	defer response.Body.Close()

	if copyErr != nil {
		outputChannel <- fmt.Sprintf("Error copying stream %v", err)
	}

	timeTaken := time.Since(startTime).Milliseconds()

	output := fmt.Sprintf("Url:%s\tLength:%d\tTime Taken: %d", url, nBytes, timeTaken)

	outputChannel <- output

}
