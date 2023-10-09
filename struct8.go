package main

type viagem struct {
	origem  string
	destino string
	data    int
	preco   []float64
}

func viagemCara(v viagem) struct {
	origem  string;
	destino string;
	data    int;
	preco   float64
} {
	var maior float64
	for _, i := range v.preco {
		if i > maior {
			maior = i
		}
	}
	viagemMaisCara := viagem{origem: v.origem, destino: v.destino, data: v.data, preco: maior}
	return viagemMaisCara
}

