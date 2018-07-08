package main

import (
	"fmt"
)

func main() {

	var notas [3]float64

	notas[0] = 6
	notas[1] = 8
	notas[2] = 10

	var soma float64 = 0

	for i := 0; i < len(notas); i++ {

		soma += notas[i]

	}

	var media float64 = soma / float64(len(notas))

	fmt.Print(media)

}
