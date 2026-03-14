package main

import (
	"fmt"
	"os"
)

//This prints the boiling point of water

func main() {
	const freezingF, boilingF = 32.0, 212.0

	fmt.Fprintf(os.Stdout, "\nThe boiling point of water is %.2F or %.2FC\n", boilingF, ftoc(boilingF))
	fmt.Fprintf(os.Stdout, "\nThe freezing point of water is %.2F or %.2FC\n", freezingF, ftoc(freezingF))
}

func ftoc(degreesInF float64) float64 {
	return (degreesInF - 32) * 5 / 9
}
