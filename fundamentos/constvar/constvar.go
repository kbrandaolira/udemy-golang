package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio float64 = 3.2

	// forma reduzida

	area := PI * math.Pow(raio, 2)

	fmt.Print("A área da circunferência é ", area, "\n")

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Print(a, b, c, d)

}
