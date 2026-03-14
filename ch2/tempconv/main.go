// This package handles the conversion between Celsius and Farenheit
package tempconv

import "fmt"

// Package tempconv performs Celsius and Fahrenheit temperature
// computation

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC  Celsius = -273.15
	BoilingPointF  Celsius = 0
	FreezingPointF Celsius = 100
)

func celsiusToF(tempC Celsius) Fahrenheit {
	return Fahrenheit((tempC * 9 / 5) + 32)
}

func fahrenheitToC(tempF Fahrenheit) Celsius {
	return Celsius(tempF-32) * 5 / 9
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f°F", f)
}

func main() {
	fmt.Printf("Absolute Zero %v\n", AbsoluteZeroC)
}
