package main

import (
	"fmt"
	"os"
)

// This is the function that executes
func main() {
	fmt.Println("Hello World")

	x := os.Args[1:]

	for x, y := range x {
		fmt.Printf("%d,%v\n", x, y)

	}

	for initialization(); ; {
		break
	}
}

func initialization() {
	fmt.Println("This is an initialization fucntion that should test the loop")
}
