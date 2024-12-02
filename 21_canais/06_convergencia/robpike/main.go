package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Vou criar dois canais com a função trabalho(), vou botar esses dois canais na função converge(), a
	// função converge() vai criar um canal, vai tacar tudo daqueles dois canais nesse um canal e vai me
	// dar esse canal novo.
	canal := converge(trabalho("maçã"), trabalho("pêra"))

	// Demonstrando 16 valores do canal, então o que estou fazendo aqui é: enquanto o meu for estiver
	// rodando vou tirar 16 vezes um valor do canal e vou botar ele na tela.
	for x := 0; x < 16; x++ {
		fmt.Println(<-canal)
	}
}

// Essa função cria um canal, cria uma go func que manda coisas para esse canal e me dá o canal devolta.
func trabalho(s string) chan string {
	canal := make(chan string)

	go func(s string, c chan string) {
		for i := 1; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			// Simulando um workload com um tempo aleatório entre nada e 1000 milisegundo (1s).
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, canal)

	return canal
}

// Essa função recebe dois canais e retorna outro canal.
func converge(x, y chan string) chan string {
	novo := make(chan string)

	go func() {
		for {
			// Aqui estou dizendo que o canal novo vai receber tudo que sair do canal x.
			novo <- <-x
		}
	}()

	go func() {
		novo <- <-y
	}()

	return novo
}

// Agora vamos pensar o que esse código está fazendo? Por que isso é uma convergência? Porque eu tenho dois
// canais criados cada um com a sua função, esses dois canais executam um trabalho que está simbolizado por
// esse tempo que esperamos, então eu tenho dois canais gerando informação e botando informação nos canais.
// Eu tenho uma função que converge esses resultados de dois canais para um só e ao final estou tirando
// informações desse um canal só.
