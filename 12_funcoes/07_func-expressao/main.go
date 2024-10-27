package main

import "fmt"

func main() {
	x := 12

	// Aqui estou atribuindo a uma variável um valor que não é uma string, não é um int, não é um struct, não é um map,
	// não é um slice, é uma função porque a função podemos utilizar ela como um valor da mesma maneira que a gente pode
	// utilizar qualquer coisa como valor.
	y := func(x int) int {
		return x * 873648
	}

	fmt.Println(x, "vezes 873648 é:", y(x))
}
