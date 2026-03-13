package main

import (
	"fmt"

	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Instead of using 'streaming' mode as in [dup2] and [dup1]
// load the entire input into memory, split it into lines
// and process it
func main() {
	counts := make(map[string]int)
	paths := os.Args[1:]
	if len(paths) == 0 {
		log.Fatal("You have not provided file paths to execute this file")
	}

	path := paths[0]
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read file with err: %v", err)
	}

	for _, line := range strings.Split(string(file), "\n") {
		counts[line]++
	}

	fmt.Printf("\nSummary\n")

	for k, v := range counts {
		if v > 1 {
			fmt.Printf("\n%s:%d\n", k, v)
		}
	}

}
