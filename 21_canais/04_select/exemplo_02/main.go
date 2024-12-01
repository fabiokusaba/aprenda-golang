package main

import "fmt"

func main() {
	canal := make(chan int)
	quit := make(chan int)
	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
}

func recebeQuit(canal chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido:", <-canal)
	}
	quit <- 0
}

func enviaPraCanal(canal chan int, quit chan int) {
	qualquercoisa := 1

	// For infinito
	for {
		// select
		select {
		case canal <- qualquercoisa:
			qualquercoisa++
		case <-quit:
			return
		}
	}
}

// O que está acontecendo aqui? Aqui eu tenho um canal quit e eu tenho um select que ele manda coisa para o
// canal e ele recebe coisas de um canal, então está aí mais uma funcionalidade do select em que ele pode
// tanto enviar quanto receber e a gente pode fazer uso dessa característica dele.
