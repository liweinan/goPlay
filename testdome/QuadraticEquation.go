package main

// https://www.testdome.com/questions/golang/quadratic-equation/85995
import (
	"fmt"
	"math"
)

func findRoots(a, b, c float64) (float64, float64) {
	// Calculate the discriminant
	d := b*b - 4*a*c

	// Calculate the square root of the discriminant
	sqrtD := math.Sqrt(d)

	// Compute the two roots using the quadratic formula
	x1 := (-b + sqrtD) / (2 * a)
	x2 := (-b - sqrtD) / (2 * a)

	// Return the roots (order doesn't matter)
	return x1, x2
}

func main() {
	x1, x2 := findRoots(2, 10, 8)
	fmt.Printf("Roots: %f, %f", x1, x2)
}
