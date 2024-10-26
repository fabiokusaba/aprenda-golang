package main

import "fmt"

// Declarando um array de inteiros com tamanho 5.
var x [5]int

var y [6]int

func main() {
	// Atribuindo valores ao array declarado.
	x[0] = 1
	x[1] = 10

	fmt.Println(x[0], x[1])
	fmt.Println(x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	// Funções que podemos utilizar com arrays.
	fmt.Println(len(x), len(y)) // Tamanho, quantidade de itens que o array tem.
}

// Em Go agrupamento de dados são as maneiras diferentes que a gente tem de organizar as informações dentro do nosso
// programa, vamos ver os quatro tipos principais: arrays, slices, maps e structs.
// O Array é um tipo que você não vai usar no dia a dia, mas o slice você vai usar no dia a dia e o Array é o bloco
// fundamental em cima do qual o slice é construído.
// O Array escolhemos o tipo que a gente quer, ele é uma estrutura de dados então ele vai ter vários valores e não um
// valor só que nem um string, int. Esses valores vão ter um tipo no nosso caso o tipo desses valores é int e
// integralmente a construção do Array é o tamanho dele portanto esse é um Array do tipo [5]int.
// O Array é uma sequência numerada de elementos de um único tipo, o número de elementos se chama length (comprimento)
// e ele nunca é negativo.
