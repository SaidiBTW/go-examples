package main

// Use the a programs command line arguments to set the values of certain
// variables throughout the program.
// We set -n cause echo to emit trailing new line
// We set -s <sep> to seperate the values with sep instead of the defult

import (
	"flag"
	"fmt"
	"strings"
)

var omitNewLine = flag.Bool("n", false, "Used to emit the trailing new line")
var seperator = flag.String("s", " ", "Seperate values with sep instead of the default")

func main() {
	flag.Parse()

	if *omitNewLine {

		fmt.Printf("%s", strings.Join(flag.Args(), *seperator))

	} else {
		fmt.Printf("%s\n", strings.Join(flag.Args(), *seperator))
	}

}
