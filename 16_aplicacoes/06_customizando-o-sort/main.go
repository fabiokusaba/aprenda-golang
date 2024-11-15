package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []carro

func (x ordenarPorPotencia) Len() int {
	return len(x)
}
func (x ordenarPorPotencia) Less(i, j int) bool {
	return x[i].potencia < x[j].potencia
}
func (x ordenarPorPotencia) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int {
	return len(x)
}
func (x ordenarPorConsumo) Less(i, j int) bool {
	return x[i].consumo > x[j].consumo
}
func (x ordenarPorConsumo) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type ordenarPorLucroProDonoDoPosto []carro

func (x ordenarPorLucroProDonoDoPosto) Len() int {
	return len(x)
}
func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool {
	return x[i].consumo < x[j].consumo
}
func (x ordenarPorLucroProDonoDoPosto) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	carros := []carro{
		carro{"Chevette", 50, 8},
		carro{"Porsche", 300, 5},
		carro{"Fusca", 20, 30},
	}

	// Slice original
	fmt.Println("Inicial:\n", carros)

	// Ordenando por potencia
	sort.Sort(ordenarPorPotencia(carros))

	// Slice ordenado
	fmt.Println("Potência:\n", carros)

	// Ordenando por consumo
	sort.Sort(ordenarPorConsumo(carros))

	// Slice ordenado
	fmt.Println("Consumo:\n", carros)

	// Ordenando por lucro para o dono do posto
	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))

	// Slice ordenado
	fmt.Println("Lucro:\n", carros)
}
