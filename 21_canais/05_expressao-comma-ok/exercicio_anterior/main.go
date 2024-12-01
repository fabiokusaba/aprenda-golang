package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	go mandaNumeros(par, impar, quit)

	receive(par, impar, quit)
}

func mandaNumeros(par, impar chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}

	close(par)
	close(impar)
	quit <- true
}

func receive(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é par.")
		case v := <-impar:
			fmt.Println("O número", v, "é ímpar.")
		// A gente colocando comma ok aqui a gente sabe dizer exatamente se deu um comma ok negativo, se a
		// gente recebeu um comma ok negativo no canal quit, bom isso não é pra acontecer porque a gente só
		// tem positivo. Se a gente recebeu um comma ok true a gente está encerrando e então a gente vai
		// retornar essa função.
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Deu zebra! Ó:", v)
			} else {
				fmt.Println("Encerrando... Recebemos:", v)
			}
			return
		}
	}
}
