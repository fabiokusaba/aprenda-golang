package main

import "fmt"

func main() {
	x := 0
	y := &x         // Armazenando o endereço em memória da variável x
	fmt.Println(*y) // Dereferencing -> saber o que tem lá dentro daquele endereço que eu tenho
	fmt.Println(&x) // Mostrando o endereço em memória onde está salvo a variável

	fmt.Println(x, y) // 0, 0x104

	*y++ // Incrementando o valor armazenado no endereço da variável x

	fmt.Println(x, y) // 1, 0x104

	// x é um int e y é um *int (ponteiro), ou seja, o ponteiro ele é o endereço. O ponteiro é uma variável
	// que armazena um endereço na memória.
	fmt.Printf("%T, %T", x, y)
}

// Todos os valores que a gente tem armazenado no computador eles ficam em um endereço na memória e a gente
// pode acessar esse endereço através do endereço, ou seja, a gente tem um número que é o endereço na memória
// e a gente pode usar esse número para acessar a informação que está lá e a gente pode pegar esse endereço e
// brincar com ele, jogar ele pra cima e pra baixo como se ele fosse uma variável.
// Um ponteiro é um objeto da linguagem de programação cujo valor se refere a outro valor salvo em outra
// localização na memória do computador utilizando esse endereço de memória, então um ponteiro faz referência
// a uma localização na memória e para obter o valor salvo naquela localização da memória a gente usa o
// procedimento chamado dereferencing.
