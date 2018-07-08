package main

import (
	"fmt"
)

type item struct {
	id         int
	quantidade int
	preco      float64
}

type pedido struct {
	id    int
	itens []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}

	return total

}

func main() {

	pedido := pedido{
		id: 1,
		itens: []item{
			item{1, 2, 5.0},
		},
	}

	fmt.Printf("Pedido %d possui o valor total igual a %.2f", pedido.id, pedido.valorTotal())

}
