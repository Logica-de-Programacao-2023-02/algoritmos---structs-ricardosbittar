package main

import (
	"fmt"
)

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func (f *Filme) AdicionarAvaliacao(avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func (f *Filme) RemoverUltimaAvaliacao() {
	if len(f.Avaliacoes) > 0 {
		f.Avaliacoes = f.Avaliacoes[:len(f.Avaliacoes)-1]
	}
}

func (f *Filme) CalcularMediaAvaliacoes() float64 {
	if len(f.Avaliacoes) == 0 {
		return 0.0
	}

	soma := 0
	for _, avaliacao := range f.Avaliacoes {
		soma += avaliacao
	}

	return float64(soma) / float64(len(f.Avaliacoes))
}

func (f *Filme) ImprimirInformacoes() {
	fmt.Printf("Título: %s\n", f.Titulo)
	fmt.Printf("Diretor: %s\n", f.Diretor)
	fmt.Printf("Ano: %d\n", f.Ano)
	fmt.Printf("Média de Avaliações: %.2f\n", f.CalcularMediaAvaliacoes())
}
