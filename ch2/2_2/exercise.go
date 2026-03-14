// Write a general-purpose unit-conversion program analogous to cf
// that reads numbers from its command-line arguments or from the standard
// input if there are no arguments and converts each number into units
// like temperature in Celsius and Farnheit, length in feet and metersm
// weight in pounds an
// ----__CF__-----
// import (
// "fmt"
// "os"
// "strconv"
// "gopl.io/ch2/tempconv"
// )
// func main() {
// for _, arg := range os.Args[1:] {
// t, err := strconv.ParseFloat(arg, 64)
// if err != nil {
// fmt.Fprintf(os.Stderr, "cf: %v\n", err)
// os.Exit(1)
// }
// f := tempconv.Fahrenheit(t)
// c := tempconv.Celsius(t)
// fmt.Printf("%s = %s, %s = %s\n",
// f, tempconv.FToC(f), c, tempconv.CToF(c))
// }
// }
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

const (
	tempConversionFToCIndex = iota
	tempConversionCToFIndex
	tempConversionCToKIndex
	tempConversionKtoCIndex

	lengthConversionFToM
	lengthConversionMtoF
)

var currentState int64 = -1

func main() {
	flag.Parse()
	fmt.Printf("GENERIC CONV PROGRAM\n--------------------")
	input := bufio.NewScanner(os.Stdin)

	shouldExit := false

	for !shouldExit {
		if !(currentState < 0) {
			switch currentState {
			case tempConversionFToCIndex:
				{
					fmt.Println("Converting from F to C")

				}
			case tempConversionCToFIndex:
			case tempConversionCToKIndex:
			case tempConversionKtoCIndex:
			case lengthConversionFToM:
			case lengthConversionMtoF:
			}
		} else {
			fmt.Println("Please Select your Input")
			fmt.Println("1. Convert Between F° to C°")
			fmt.Println("2. Convert Between C° and F°")
			fmt.Println("3. Convert Between C° to K")
			fmt.Println("4. Convert Between K to C°")
			fmt.Println("5. Convert Between Feet to Metres")
			fmt.Println("5. Convert Between Metres to Feet")
			fmt.Print("Input: ")
			input, err := strconv.ParseInt(input.Text(), 10, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Invalid Input: %v", err)
			}
			currentState = input
		}
	}

}
