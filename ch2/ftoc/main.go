// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g℉ = %g℃\n", freezingF, fToc(freezingF)) //"32℉ = 0℃"
	fmt.Printf("%g℉ = %g℃", boilingF, fToc(boilingF))     //"212℉ = 100℃"
}

func fToc(f float64) float64 {
	return (f - 32) * 5 / 9
}
