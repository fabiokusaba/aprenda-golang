package main

import "fmt"

func main() {
	x := 10
	if !(x > 100) {
		fmt.Println("Hello, playground")
	}

	if y := 10; !(y > 100) {
		fmt.Println("Hello, playground")
	}
}

// O if é relativamente simples comparado ao for, a gente tem uma condição aqui se essa condição for true o que estiver dentro
// do codeblock será executado, se for false não executa.
// Operador de negação "!" -> basicamente o que ele faz é inverter o resultado de qualquer expressão de comparação, ou seja, ele é
// o não. Tudo o que é verdadeiro ou falso ele diz não, ou seja, o contrário.
// if declaração de inicialização -> posso dar uma condição inicial antes do if rodar.
