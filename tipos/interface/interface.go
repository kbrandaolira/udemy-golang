package main

import "fmt"

type voador interface {
	voar()
}

type calopsita struct {
	nome string
}

type periquito struct {
	nome string
}

func (c calopsita) voar() {
	fmt.Printf("Calopsita %s está voando!", c.nome)
}

func (p periquito) voar() {
	fmt.Printf("Periquito %s está voando!", p.nome)
}

func voar(v voador) {
	v.voar()
}

func main() {

	var c voador = calopsita{"Caca"}

	voar(c)

}
