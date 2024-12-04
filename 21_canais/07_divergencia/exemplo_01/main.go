package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)

	go manda(20, canal1)
	go outra(canal1, canal2)

	for v := range canal2 {
		fmt.Println(v)
	}
}

func manda(n int, canal chan int) {
	for i := 0; i < n; i++ {
		canal <- i
	}
	close(canal)
}

func outra(canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for v := range canal1 {
		wg.Add(1)
		go func(x int) {
			canal2 <- trabalho(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}

// Divergência é o contrário de convergência, ou seja, na convergência a gente fez exemplos em que a gente tinha várias
// coisas que iam a um ponto só, agora a gente vai ter poucas coisas que vão a muitos pontos.
// Vamos entender o que fizemos aqui: eu tenho dois canais, ou seja, o meu programa vai começar com um canal e vai
// terminar com um canal, mandamos x números para o primeiro canal, depois criamos uma outra função que pega cada
// informação que eu botei ali, cada item que foi no primeiro canal e manda uma goroutine para cada um, ou seja, nesse
// momento eu vou ter vinte goroutines extras para fazer esse trabalho em execução, ou seja, vou ter vinte itens no
// canal1 e para cada um deles eu vou chamar uma goroutine, quando essas goroutines forem executadas, ou seja, eu vou
// executar a função trabalho() vinte vezes, quando essas goroutines terminarem a sua execução elas vão mandar o
// resultado do trabalho delas para o canal2 então o que aconteceu? Eu tinha um canal eu espalhei esse trabalho() entre
// várias goroutines que poderiam ser milhões de goroutines depois eu convergi esse valor denovo para um canal, então
// aqui temos um exemplo duplo de divergência e convergência.
