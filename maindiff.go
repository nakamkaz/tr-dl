package main

import (
	"log"
	"math"
)

var (
	pi float64 = math.Pi
)

func main() {
	dSin := Dydx(math.Sin) // ~~ Cos(x)
	dCos := Dydx(math.Cos) // ~~ -Sin(x)
	dExp := Dydx(math.Exp)
	for _, f := range []float64{0.0, 0.5, 1.0, 1.5, 2.0} {
		log.Print("# Cos'(", f, "*pi): ", dCos(f*pi), " ~~ -Sin(", f, "*pi): ", -1.0*math.Sin(f*pi))
		log.Print("* Sin'(", f, "*pi): ", dSin(f*pi), " ~~  Cos(", f, "*pi): ", math.Cos(f*pi))
		log.Print(dExp(f), math.Exp(f))
	}
}

