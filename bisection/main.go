package main

import (
	"fmt"
	"math"
)

func main() {
	x0 := 0.0
	x1 := 1.0
	iterations := 0
	for math.Abs(x0-x1) > 1e-5 {
		x1, x0 = bisection(x0, x1)
		iterations++
		fmt.Printf("Iteration %d: x0 = %.6f, x1 = %.6f\n", iterations, x0, x1)
	}

	fmt.Printf("Total iterations: %d\nx0: %.6f\nx1: %.6f\n ", iterations, x0, x1)
}

func f(val float64) float64 {
	return math.Pow(val, 3) - (7 * math.Pow(val, 2)) + (14 * val) - 6
}

func bisection(x0 float64, x1 float64) (float64, float64) {
	if f(x0)*f(x1) < 0 {
		half := (x0 + x1) / 2
		if f(half)*f(x0) < 0 {
			return half, x0
		} else if f(half)*f(x1) < 0 {
			return half, x1
		}
	}
	return x1, x0
}
