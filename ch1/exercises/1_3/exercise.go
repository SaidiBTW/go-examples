package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Experiment to measure the differnce in running time between our potentially inefficient versions and the one the uses strings join

func Version1() {
	startTime := time.Now()
	args := os.Args

	var s string

	sep := " "

	for i := 0; i < len(args); i++ {
		s += s + args[i] + sep
	}
	fmt.Printf("\nInefficient Version String\n: %s", s)

	fmt.Printf("\nThe function took %d Milliseconds\n", time.Since(startTime).Milliseconds())
}

func Version2() {
	startTime := time.Now()
	args := os.Args

	s := strings.Join(args, " ")

	fmt.Printf("\nEfficient Version String: %s\n", s)

	fmt.Printf("\nThe function took %d Milliseconds\n", time.Since(startTime).Milliseconds())

}

func main() {
	Version1()
	Version2()

}
