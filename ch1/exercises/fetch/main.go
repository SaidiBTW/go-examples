package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Print contents found at the url

func main() {
	//Get the value from Args
	url := os.Args[1:][0]

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error fetching url: %s\nError: %v", url, err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "There was an error reading response: %s\nError: %v", url, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Response: %s", b)

}
