package main

import "fmt"

func main() {
	x := 0
	estarecebeovalor(x)
	estarecebeumponteiro(&x)

	fmt.Println(x)
}

func estarecebeovalor(x int) {
	x++ // Incrementa uma cópia do valor
	fmt.Println("Na função:", x)
}

func estarecebeumponteiro(x *int) {
	*x++ // Pegando o valor que estava no endereço da variável x e acrescentando +1
	fmt.Println("Na função:", *x)
}

// A utilidade prática dos ponteiros -> ponteiros servem para duas coisas: uma é pra quando estou lidando com
// quantidades bem grandes de dados e eu não quero ficar copiando esse valor pra lá e pra cá, eu quero deixar
// ele em um lugar na memória e todos que tiverem que mexer com esse valor vai lá naquele lugar e acessa direto
// então, por exemplo em Go tudo é "pass by value" quando uma função recebe um valor eu tenho que copiar esse
// valor pra função daí a função executa alguma coisa em cima daquilo e essa cópia tem um custo de performance
// o que eu posso fazer? Ao invés de passar pra função o valor eu passo simplesmente um ponteiro pro valor, então
// ao invés de eu copiar o valor inteiro eu só digo isso aqui que você quer está lá naquele endereço de memória
// aponta pra lá quando você for usar e aí eu não tenho que fazer a cópia.
// Outra utilidade é quando a gente quer mudar o valor na posição original.
