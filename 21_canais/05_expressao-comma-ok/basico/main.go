package main

import "fmt"

func main() {
	canal := make(chan int)

	go func() {
		canal <- 42
		close(canal)
	}()

	v, ok := <-canal

	// Recebi um valor (v) e o ok está me dizendo que isso é um valor que recebi mesmo, é um valor real do
	// canal, não é simplesmente um valor padrão, então 42 é um número de verdade.
	fmt.Println(v, ok)

	// Se eu repetir o código novamente agora eu recebi o valor 0 e false, então se eu não tivesse o comma
	// ok ali como é que eu ia saber se esse 0 é um número que realmente o programa passou.
	// Mudando o nosso exemplo para 0, perceba que recebemos o valor 0 duas vezes, mas uma vez realmente o
	// valor que estava naquele canal era o 0, na segunda vez não, na segunda vez eu recebi 0 porque não
	// tinha nada lá, então foi um valor padrão.
	// Para eu saber essa diferença eu vou usar o comma ok.
	v, ok = <-canal
	fmt.Println(v, ok)
}

// A gente já falou da expressão comma ok no contexto de maps e agora vamos ver como ela funciona no contexto
// de canais.
