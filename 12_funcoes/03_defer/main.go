package main

import "fmt"

func main() {
	// Quando temos mais de um defer aquele que entrou primeiro na lista será executado por último.
	defer fmt.Println("1") // defer -> deixa para última hora.
	defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")

}

// A funcionalidade defer pode ser explicado como deixar para última hora, ou seja, você vai dar uma instrução e em
// algum momento vai ter um defer na frente isso vai significar que você vai executar aquele bloco de código e quando
// chegar na última hora vai executar o defer.
// A utilidade disso é deixar o seu códgo modular, ou seja, suponhamos que você tenha uma função que abre um arquivo
// sempre que você abre um arquivo você vai ficar com um espaço de memória alocado para as informações referentes a
// esse arquivo, então sempre depois de terminar de usar um arquivo você deve fechá-lo.
