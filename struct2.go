package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}
type endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}

func imprimeEnd(e endereco) {
	fmt.Print(e.rua)
	fmt.Print(e.numero)
	fmt.Print(e.cidade)
	fmt.Print(e.estado)
}
