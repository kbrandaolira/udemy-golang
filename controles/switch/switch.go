package main

import (
	"fmt"
)

func main() {

	fmt.Print(imprimirConceito(10))
	fmt.Print(imprimirConceito(5))
	fmt.Print(imprimirConceito(4))
	fmt.Print(imprimirConceito(1))

}

func imprimirConceito(nota float64) string {

	switch nota {
	case 10, 9, 8, 7, 6:
		return "A"
	case 5:
		fallthrough
	case 4:
		return "B"
	case 3, 2, 1:
		return "C"
	default:
		return "Z"
	}

}
