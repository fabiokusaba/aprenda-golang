package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	funcoes := 5

	go manda(100, canal1)
	go outra(funcoes, canal1, canal2)

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

func outra(funcoes int, canal1, canal2 chan int) {
	var wg sync.WaitGroup

	for i := 0; i < funcoes; i++ {
		wg.Add(1)
		go func() {
			for v := range canal1 {
				canal2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(canal2)
}

func trabalho(n int) int {
	time.Sleep(time.Millisecond * 1000)
	return n
}

// Perceba que aqui temos pedaços de cinco em cinco, ou seja, cada trabalho() leva um segundo para executar a gente tem
// cinco funções rodando de cada vez, então a cada segunda a gente vai ter a resposta de cinco funções e é isso que
// vemos no terminal, isso prova que eu tenho cinco goroutines executando passando por todos aqueles valores até acabar
// então se a gente for ver aquela linha do tempo como é que ficou: eu tinha um stream esse stream virou 5 goroutines
// essas goroutines foram trabalhando, trabalhando, trabalhando... A medida em que elas trabalhavam elas pegavam esses
// cinco valores e colocavam em um ponto só e aí elas encerraram, mas o interessante é que eu sempre tinha cinco coisas
// sendo executadas e esses cinco valores eram colocados em um único canal que recebia esses valores e botava na tela.
// Esses foram os nossos exemplos de divergência, ou seja, eu tenho um stream, abro esse stream em milhares de threads
// diferentes e essas threads depois, já não é divergência seria convergência, manda o resultado pra um lugar só e esse
// um lugar só coloca as coisas na tela pra gente.
