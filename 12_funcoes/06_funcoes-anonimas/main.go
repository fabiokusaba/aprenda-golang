package main

import "fmt"

func main() {
	x := 387

	// Aqui temos uma função que não é igual as outras funções, declarei a função e chamei ao mesmo tempo.
	// Isso é uma função anônima e ela não precisa de um nome.
	// Isso serve para quando você quer uma função descartável em que você quer usar uma vez e deu e também para quando
	// vamos trabalhar com goroutines.
	func(x int) {
		fmt.Println(x, "vezes 873648 é:")
		fmt.Println(x * 873648)
	}(x)

}
