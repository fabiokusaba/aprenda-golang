package main

import "fmt"

func main() {
	x := 0
	for x <= 20 {
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}

// O que esse exemplo faz é mostrar todos os números pares.
// Se o módulo do número dividido por dois for zero significa que ele é par, se não for significa que ele é ímpar.
// Lembra quando falamos do break, o break ele pega o loop inteiro e joga fora ele quebrou o loop inteiro.
// O continue ele quebra a iteração atual, então o que estamos fazendo aqui é que quando o número não for par eu paro a iteração
// do loop e vou para a próxima.
