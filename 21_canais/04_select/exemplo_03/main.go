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
	// Aqui vou mandar números pares para um canal, ímpares para outro canal e fechar e mandar um quit
	total := 100
	for i := 0; i < total; i++ {
		// Se i for divisível por 2 vou mandar o valor para o canal par
		if i%2 == 0 {
			par <- i
		} else {
			// Caso contrário, vou mandar o valor para o canal impar
			impar <- i
		}
	}
	// Ao final vou dar um close e para o canal quit vou mandar um true
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
		case <-quit:
			return
		}
	}
}
