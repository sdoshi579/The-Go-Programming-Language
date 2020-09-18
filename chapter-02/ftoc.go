// Prints two Fahrenheit-to-Celsius conversions
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0

	// %g removes the trailing zeros after decimal
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))

	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float32) float32 {
	return (f - 32) * 5 / 9
}

/*
%f in printf prints 32.000000°F = 0.000000°C
%g in printf prints 32°F = 0°C
*/
