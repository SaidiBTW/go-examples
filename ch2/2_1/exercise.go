// In this package ADD types constants and function for processing temperatures in the
// Kelvin Scale

package main

import "fmt"

type Kelvin float64
type Celsius float64
type Farenheigh float64

const (
	AbsoluteZeroInK     = 0
	WaterFreezingPointK = 273.15
	WaterBoilingPointK  = 373.15
)

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2fK", k)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k + 273.15)
}
