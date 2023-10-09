package main

type animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

func (a *animal) trocarSom(novoSom string) {
	a.som = novoSom
}
