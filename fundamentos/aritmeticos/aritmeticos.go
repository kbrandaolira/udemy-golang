package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3
	b := 2

	fmt.Print("Soma", a+b, "\n")
	fmt.Print("Subtração", a-b)
	fmt.Print("Divisão", a/b)
	fmt.Print("Multiplicação", a*b)
	fmt.Print("Módulo", a%2)

	fmt.Print("Maior? ", math.Max(float64(a), float64(b)))
}
