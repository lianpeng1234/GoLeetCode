package main

import (
	"math"
)

// 2 9
// 2 6
// 2 2
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = int(math.Abs(float64(n)))
		return myPowDiGui(x, n)
	} else if n == 0 {
		return 1
	} else {
		return myPowDiGui(x, n)
	}

}

func myPowDiGui(x float64, n int) float64 {
	if n == 1 {
		return x
	}
	if n == 0 {
		return 1
	}

	if n%2 == 1 {
		value := myPowDiGui(x, n/2)
		return value * value * x
	} else {
		value := myPowDiGui(x, n/2)
		return value * value
	}

}
