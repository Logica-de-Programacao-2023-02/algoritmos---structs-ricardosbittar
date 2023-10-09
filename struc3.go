package main

type Triangulo struct {
	base   int
	altura int
}

func areaTri(t Triangulo) float64 {
	area := t.base * t.altura
	return float64(area / 2)
}
