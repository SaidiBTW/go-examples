package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// Modyfy fetch to add the prefix http:// to each argument URL
// if it is missing. You might want to use strings.HasPrefix
func main() {
	urls := os.Args[1:]
	if len(urls) == 0 {
		log.Fatal("Please pass a url in the program arguments")
	}

	for _, url := range urls {
		var finalUrlString string = url
		// Checking for prefix in url
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			finalUrlString = "http://" + finalUrlString
		}
		response, err := http.Get(finalUrlString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError fetching url \"%s\" : Err: %v", finalUrlString, err)
			// To make the program a bit more fault tolerant
			// We use continue to skip further processing
			// and process other elements in the loop
			continue
		}

		_, copyError := io.Copy(os.Stdout, response.Body)
		defer response.Body.Close()

		if copyError != nil {
			fmt.Fprintf(os.Stderr, "\nError copying output to os.StdOut: Err: %v", copyError)
			// To make the program a bit more fault tolerant
			// We use continue to skip further processing
			// and process other elements in the loop
			continue
		}

	}

}
