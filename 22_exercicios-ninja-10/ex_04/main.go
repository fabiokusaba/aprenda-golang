package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	// Aqui sabemos que vai ter que ser uma go func igual o exercício anterior porque eu não posso ter uma função que
	// recebe e uma função que envia rodando ao mesmo tempo
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()

	return c
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Printf("received %d\n", v)
		case <-q:
			return
		}
	}
}
