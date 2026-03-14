package main

import (
	"fmt"
	"os"
)

//This prints the boiling point of water

const boilingF = 212.0

func main() {

	fmt.Fprintf(os.Stdout, "\nThe boiling point of water is %.2F or %.2FC\n", boilingF, ftoc(212))
}

func ftoc(degreesInF float64) float64 {
	return (degreesInF - 32) * 5 / 9
}
