package main

import "fmt"

func main() {
	a := make(chan int)
	b := make(chan int)
	x := 500

	// Esse x aqui vai agir dentro do codeblock
	go func(x int) {
		for i := 0; i < x; i++ {
			// Enviando o valor de i para o canal a
			a <- i
		}
	}(x / 2) // Esse x aqui está no escopo da função main

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)

	// For loop para iterar sobre os valores
	for i := 0; i < x; i++ {
		// select: aprendemos que ele é igual a um switch com cases e eu consigo usar esse select para
		// receber valores de vários canais em uma mesma localização, então não preciso receber
		// especificamente de uma coisa separada para cada canal, eu posso ter uma função que vai receber
		// coisas de vários canais
		select {
		case v := <-a:
			fmt.Println("Canal A:", v)
		case v := <-b:
			fmt.Println("Canal B:", v)
		}
	}
}

// Select é como um switch só que para canais e ele não é sequencial.
// Como vimos na primeira aula sobre canais ele bloqueia até ele receber um valor que bate com os casos
// que a gente especificou nele, quando ele recebe ele executa.
// Se eu tiver mais de um caso válido ao mesmo tempo, então ele vai escolher um aleatório e eu não vou
// ter controle sobre o qual ele vai escolher.
