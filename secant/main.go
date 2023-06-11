package main

import (
	"fmt"
	"math"
)

func main() {
	cur := 1.0
	prev := 0.0
	iterations := 0
	for math.Abs(cur-prev) > 1e-5 {
		cur, prev = secant(cur, prev)
		iterations++
		fmt.Printf("Iteration %d: cur = %.6f, prev = %.6f\n", iterations, cur, prev)
	}

	fmt.Printf("Total iterations: %d\nThe root: %.6f", iterations, cur)
}

func f(val float64) float64 {
	return math.Pow(math.E, -val) - val
}

func secant(cur float64, prev float64) (float64, float64) {
	return cur - ((f(cur) * (prev - cur)) / (f(prev) - f(cur))), cur
}
