package main

import (
	"fmt"
	"math"
)

// The function x^(1/5) 's reverse
func f(x float64) float64 {
	return math.Pow(x, 5) - 16
}

// The function x^(1/5) 's derrivative's reverse
func df(x float64) float64 {
	return 5 * math.Pow(x, 4)
}

// Implementation of Newton Raphson method
func newtonRaphson(val float64) (float64, int) {
	x := val
	count := 1
	for i := 0; i < 100; i++ {
		fx := f(x)
		dfx := df(x)
		x = x - fx/dfx
		// Stop if the difference is smaller than 10^-5
		if math.Abs(fx/dfx) < 1e-5 {
			break
		}
		// Increment the iteration counter
		count++
	}
	return x, count
}

// Implementation of Bisection method
func bisection(a, b float64) (float64, int) {
	fa := f(a)
	fb := f(b)
	count := 1
	if fa*fb > 0 {
		panic("f(a) and f(b) must have opposite signs")
	}
	for i := 0; i < 100; i++ {
		c := (a + b) / 2
		fc := f(c)
		// Stop if the difference is smaller than 10^-5
		if fc == 0 || (b-a)/2 < 1e-5 {
			return c, count
		}
		if fa*fc > 0 {
			a = c
		} else {
			b = c
		}
		// Increment the iteration counter
		count++
	}
	panic("method failed to converge")
}

func main() {
	x0 := 1.5 // initial guess for Newton-Raphson method
	x, count := newtonRaphson(x0)
	fmt.Printf("Newton-Raphson result: %v, iteration count: %d\n", x, count)

	// The result should be between these
	a := 1.0 // lower bound for bisection method 1^5=1
	b := 2.0 // upper bound for bisection method 2^5=32
	x, count = bisection(a, b)
	fmt.Printf("Bisection result: %v, iteration count: %d\n", x, count)
}
