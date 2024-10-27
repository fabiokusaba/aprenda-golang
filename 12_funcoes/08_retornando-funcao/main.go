package main

import "fmt"

func main() {
	// Eu tenho uma função que retorna uma função e esse retorno se tornou o valor da variável x.
	x := retornaumafuncao()

	// Passando o argumento 3 para essa função retornou o valor de 30.
	y := x(3)

	// Esse valor retornado foi impresso aqui.
	fmt.Println(y)

	// Uma outra forma de execução.
	fmt.Println(retornaumafuncao()(3))
}

// Aqui eu tenho uma função que retorna uma função e essa função que ela está retornando vai retornar um int.
// Funções dentro de funções.
func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
