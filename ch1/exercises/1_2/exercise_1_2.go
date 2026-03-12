package main

import (
	"fmt"
	"os"
)

//Modify the exho program to print the index and value of each of its arguments, one per line

func main() {
	args := os.Args

	for i := 0; i < len(args); i++ {
		fmt.Printf("%d: %s\n", i, args[i])
	}
}
