package main

func Dydx(pf func(x float64) float64) func(x float64) float64 {
	var h float64 = 0.0001
	return func(x float64) float64 { return (pf(x+h) - pf(x-h)) / 2*h }
}
