package main

import (
	"fmt"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {

	produto1 := produto{
		"Lápis", 2.00, 0.05,
	}

	fmt.Println(produto1, produto1.precoComDesconto())

}
