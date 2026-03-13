package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Modify the fetch to also print the HTTP
// status code, found in resp.Status
func main() {
	urls := os.Args[1:]

	if len(urls) == 0 {
		log.Fatal("Please pass urls to the program")
	}

	for _, url := range urls {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError fetching url %s, Err %v", url, err)
			continue
		}

		statusCode := response.StatusCode

		fmt.Printf("\n Status Code : %d \n", statusCode)

		_, copyError := io.Copy(os.Stdout, response.Body)
		response.Body.Close()

		if copyError != nil {
			fmt.Fprintf(os.Stderr, "\nError copying result to os.Stdout %v", err)
			continue
		}
	}
}
