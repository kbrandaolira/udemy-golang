package main

import (
	"fmt"
)

func main() {

	imprimirResultado(7.5)
	imprimirResultado(5.9)

}

func imprimirResultado(nota float64) {

	if nota >= 7.0 {

		fmt.Print("Aprovado!")

	} else {

		fmt.Print("Reprovado!")

	}

}
