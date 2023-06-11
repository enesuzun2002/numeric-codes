package main

import (
	"fmt"
	"math"
)

func main() {
	/* var a float64
	var b float64
	a = 1.5
	b = 1.5

	for i := 0; i < 3; i++ {
		a = f1(a)
		b = f2(b)
		fmt.Printf("New a value: %f \nNew b value: %f \n", a, b)
	}

	fmt.Printf("New a value: %f \nNew b value: %f \n", a, b) */
	a := 0.5
	for i := 0; i < 10; i++ {
		a = iteration(a)
		fmt.Printf("New a value: %f \n", a)
	}
}

func f(val float64) float64 {
	return math.Pow(val, 3) + 4*math.Pow(val, 2) - 10
}

func f_derrivative(val float64) float64 {
	return 3*math.Pow(val, 2) + 8*val
}

func f1(val float64) float64 {
	return 10/math.Pow(val, 2) - 4
}

func f2(val float64) float64 {
	return val - (f(val) / f_derrivative(val))
}

func iteration(val float64) float64 {
	return 3 / (4 + math.Pow(val, 2))
}
