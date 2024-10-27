package main

import "fmt"

func main() {
	// Crio um slice com esses valores.
	primeiroslice := []int{1, 2, 3, 4, 5}

	// Demonstro esse slice.
	fmt.Println(primeiroslice)

	// Crio um segundo slice que vai ser uma fatia do primeiro slice -> 1, 2, 5
	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	// O resultado do segundo slice.
	fmt.Println(segundoslice)

	// Quando mostro denovo o primeiro slice ele está diferente porque esse append foi no array que existe por baixo
	// aqui, no array subjacente, ele fez um refatiamento daquele array pra fazer com que esses valores que a gente
	// escolheu ficassem visíveis, ou seja, o segundo slice tem esse array subjacente, mas ele tem um len de 3 itens
	// ou seja, estou usando o mesmo array subjacente da outra variável só que como eu tenho 3 elementos os elementos
	// que vem depois desses 3 eu não estou vendo, mas eu usei o mesmo array subjacente que usava na outra variável.
	fmt.Println(primeiroslice)
}

// Quando uso o append o array que está por baixo do slice vai ser refatiado pra acomodar os novos elementos, se não der
// certo um novo array vai ser criado, mas se der certo a gente vai simplesmente refatiar o array que está por baixo.
