package main

type circulo struct {
	raio float64
}

func areaCirc(c circulo) float64 {
	pi := 3.14
	area := pi * c.raio * c.raio
	return area
}
