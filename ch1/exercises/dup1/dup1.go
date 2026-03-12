package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 print the text of each line that appears more than
// once in std input, preceded by its count.
func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		value := input.Text()

		counts[value]++

		currentValue := counts[value]
		if currentValue > 1 {
			fmt.Printf("\n This is a duplicate value: %d times\n", currentValue)
		}

	}

	fmt.Printf("\n\nSummary\n")
	for k, v := range counts {
		fmt.Printf("\n%s \t\t %d\n", k, v)
	}
}
