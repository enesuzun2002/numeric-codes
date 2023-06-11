package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	a = 1
	b = 2
	error := math.Pow10(-5) // For finding the correct root value
	root := bisection(a, b, error)

	for math.Abs(a-b) > error {
		a, b = calculate(a, b)
		fmt.Printf("New a value: %f \nNew b value: %f \n", a, b)
	}
	fmt.Printf("Final a value: %f \nFinal b value: %f \n", a, b)
	fmt.Printf("Root value: %f \nAbsolute error for b: %f\nAbsolute error for a: %f\nRelative error for b: %f\nRelative error for a: %f\n", root, math.Abs(b-root), math.Abs(a-root), (math.Abs(b-root) / root), (math.Abs(a-root) / root))

}

func calculate(a float64, b float64) (float64, float64) {
	return RegulaFalsi(a, b), NewtonRaphson(b)
}

func RegulaFalsi(a float64, b float64) float64 {
	return (a*f(b) - b*f(a)) / (f(b) - f(a))
}

func NewtonRaphson(val float64) float64 {
	return (val - (f(val) / f_derrivative(val)))
}

func f(val float64) float64 {
	return (val*math.Cos(val) - math.Log(val))
}

func f_derrivative(val float64) float64 {
	return (math.Cos(val) - (val * math.Sin(val)) - (1 / val))
}

func bisection(a float64, b float64, tol float64) float64 {
	for math.Abs(b-a) > tol {
		c := (a + b) / 2.0
		if f(c) == 0.0 {
			return c
		} else if f(c)*f(a) < 0 {
			b = c
		} else {
			a = c
		}
	}
	return (a + b) / 2.0
}
