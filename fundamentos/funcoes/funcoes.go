package main

import "fmt"

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Print("Valor ", valor)
}

func main() {

	imprimir(somar(4, 3))

}
