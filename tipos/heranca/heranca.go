package main

import "fmt"

type pessoa struct {
	id   int
	nome string
}

type pessoaFisica struct {
	pessoa
	cpf string
}

func main() {
	pf := pessoaFisica{}

	pf.id = 1
	pf.nome = "Kayan"
	pf.cpf = "13738659706"

	fmt.Printf("A pessoa física %s possui o CPF %s", pf.nome, pf.cpf)

}
