package main

import "fmt"

func main() {
	// Criando um canal
	c := make(chan int)

	// Chamando a função meuloop() como uma go func
	go meuloop(10, c)

	prints(c)
}

// Aqui na minha função eu recebo como parâmetro um canal send onde vou enviar valores a ele e uma variável
// do tipo inteiro que vai ser o meu total de valores.
func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		// Enviando para o meu canal o valor de i
		s <- i
	}
	// Para resolver o nosso erro/problema precisamos usar o close(), com isso estamos dizendo que não iremos
	// mais enviar valores então podemos encerrar esse canal.
	close(s)
}

func prints(r <-chan int) {
	// Utilizando um range para ler os valores, aqui estou recebendo informações do canal então ele passou
	// por esses dez valores da minha função meuloop() e quando chegou do 11º ele ficou esperando porque
	// ninguém avisou que não iria ter mais valores, quando ele fica esperando sem ter o que fazer a gente
	// tem esse erro: " all goroutines are asleep - deadlock!"
	for v := range r {
		fmt.Println("Recebido do canal:", v)
	}
}
