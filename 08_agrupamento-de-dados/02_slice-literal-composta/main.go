package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 5)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])

	// Aqui provavelmente teremos um erro, esse erro "index out of range" ele acontece pela razão que um slice é feito
	// de arrays, então cada vez que você atribui um valor novo para um slice basicamente o que acontece é que o
	// compilador joga aquele slice fora, cria um array novo e com esse array novo ele coloca aquele valor a mais, ou
	// seja, o tipo "subjacente" de um slice é um array então eu tenho um slice com cinco elementos e eu quero botar uma
	// coisa no índice 20 dele, mas ele nem tem o índice 20, porque o array subjacente dele não tem tantos elementos
	// assim.
	slice[20] = 1
	fmt.Println(slice[20])
}

// Um tipo de dados composto é qualquer tipo de dados que pode ser construído em um programa utilizando os tipos de
// dados primitivos da linguagem de programação e outros tipos de dados compostos. Ou seja, eu vou fazer tipos de dados
// compostos, estrutura de dados utilizando as primitivas da linguagem, por exemplo a gente viu int, string, bool, então
// a gente vai ter aqui no caso falando de slices um slice é um array e array é um conjunto então a gente vai ter
// conjuntos de dados que podem ser int, float, string e por aí vai, ou seja, a gente escolhe um tipo e faz uma coleção
// de dados desse tipo.
// Composite literals são a maneira de criar dados compostos em Go.
// A diferença de um array e um slice na declaração é que o array vai ter um número finito de elementos, um número pré
// especificado e o slice não vai ter, o slice pode ter quantos elementos você quiser.
// O slice eu consigo mudar o tamanho dele, coisa que eu não consigo fazer com array, o array ele é fixo.
