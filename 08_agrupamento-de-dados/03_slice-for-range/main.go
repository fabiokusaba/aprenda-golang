package main

import "fmt"

func main() {
	// Declarando um slice e colocando valores dentro dele, ele é criado com o valor fixo subjacente que no caso seria
	// o array que está por baixo desse slice.
	slice := []string{"banana", "maçã", "jaca", "pêssego"}

	// Com a função range a gente consegue passar por todos os valores.
	for indice, valor := range slice {
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}

	// A gente cria um slice o slice é criado utilizando um array subjacente, esse array tem um tamanho fixo o slice não
	// tem tamanho fixo, portanto se eu tentar adicionar um outro elemento da seguinte forma ele não deixa "index out of
	// range", ou seja, o índice está fora da extensão do slice.
	//slice[4] = "melancia"

	// Mas, se eu utilizar a função append que ela é própria para slices aí eu consigo mudar o tamanho do slice sem
	// problema nenhum.
	slice = append(slice, "melancia")

	// Posso pegar apenas o índice e jogar fora o valor.
	for indice, _ := range slice {
		fmt.Println("Esse slice tem", indice, "elementos.")
	}

	// Como também posso jogar fora o índice e ficar apenas com o valor.
	for _, valor := range slice {
		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
	}

	slice2 := []int{20, 21, 22, 23, 1, 13}
	total := 0

	for _, valor := range slice2 {
		total += valor // mesma coisa que -> total = total + valor
	}

	fmt.Println("O valor total é:", total)
}

// A função range -> range significa alcance, faixa, extensão, então o que a função range faz é atravessar toda extensão
// do slice e mostra pra gente tudo o que tem lá.
