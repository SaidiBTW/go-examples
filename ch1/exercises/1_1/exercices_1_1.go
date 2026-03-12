package main

import (
	"fmt"
	"os"
	"strings"
)

//Create a program that prints all args including the command that invoked it os.Args[0]

func main() {
	var s string
	s = strings.Join(os.Args, " ")
	fmt.Printf("%s\n", s)
}
