package main

import (
	"fmt"
)

func main() {
	// Area of the triangle: p = .5 * a * h
	p := triangleArea(12.3, 16.7)

	fmt.Printf("triangle area: %f\n", p)
	fmt.Printf("triangle area: %f\n", triangleArea(12.3, 160.7))
}

func triangleArea(a, h float64) float64 {
	return .5 * a * h
}
