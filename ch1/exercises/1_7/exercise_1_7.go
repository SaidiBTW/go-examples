package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// The function call io.Copy(dst,src) reads from src and writes
// to dst. Use it instrad of ioutil.Readll(all to copy the response body
// to os.Stdout without requiring a buffer large enough to hold the
// entire stream . Be sure to check the error result of io.Copy
func main() {
	// Get the URL from the os.Args
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Please provide your URL in the program args")
	}
	url := args[0]

	response, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching url %s, Err: %v", url, err)
		os.Exit(1)
	}

	_, copyErr := io.Copy(os.Stdout, response.Body)
	defer response.Body.Close()

	if copyErr != nil {
		fmt.Fprintf(os.Stderr, "Err copying response stream %v", err)
		os.Exit(1)
	}

}
