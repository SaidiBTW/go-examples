package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup 2 print the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files
// in the environmental variables

func main() {
	counts := make(map[string]int)
	paths := os.Args[1:]

	if len(paths) == 0 {

		fmt.Printf("\nInput your line: ")
		countLines(os.Stdin, counts)
	} else {

		for i := 0; i < len(paths); i++ {
			path := paths[i]

			file, err := os.Open(path)

			if err != nil {
				fmt.Fprintf(os.Stderr, "File failed to open due to error: %v", err)
			}
			countLines(file, counts)

		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		value := input.Text()
		counts[value]++
	}

	fmt.Printf("\nSummary\n")

	for k, v := range counts {
		fmt.Printf("\n%s:%d\n", k, v)
	}
}
