package main

import "errors"

type funcionario struct {
	nome    string
	salario float64
	idade   int
}

func aumentarSalario(f funcionario) float64 {
	f.salario = f.salario * 1.10
	return f.salario
}
func diminuirSalario(f funcionario) float64 {
	f.salario = f.salario * 0.9
	return f.salario
}
func tempoDeTrab(f funcionario) (int, error) {
	if f.idade < 18 {
		return 0, errors.New("Idade inválida")
	}
	return f.idade - 18, nil
}
