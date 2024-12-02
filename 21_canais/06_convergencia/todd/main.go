package main

import (
	"fmt"
	"sync"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go envia(par, impar)
	go recebe(par, impar, converge)

	for v := range converge {
		fmt.Println("Valor recebido:", v)
	}
}

func envia(p, i chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(c)
}

// Convergência é um padrão de programação, uma maneira de fazer as coisas que envolve a gente pegar dados
// de vários pontos diferentes e convergir eles em menos pontos ou um ponto só.
// Nesse nosso exemplo eu tinha dois canais, um canal com x informações e outro canal com x informações, eu
// usei uma função que converge tudo isso num canal só, então de dois canais eu convergi para um único canal
// e daí eu implementei uma funcionalidade em cima desse um canal só que estava ali e coloquei tudo isso na
// tela.
// Então, esse é um exemplo simples do que seria convergência.
