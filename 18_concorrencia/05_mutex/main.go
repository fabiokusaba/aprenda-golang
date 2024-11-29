package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 1000

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	// Tendo um valor do tipo Mutex eu ganhei dois métodos: Lock e Unlock
	var mu sync.Mutex

	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			// Antes de começar isso aqui vai ter um Lock
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			// E a hora que eu terminar isso aqui vai ter um Unlock
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)
}

// Mutex: exclusão mútua, o que isso faz? Ele tranca uma variável ou um trecho de código e faz com que somente
// uma linha de execução, um thread, uma goroutine possa mexer com aqueles valores num dado momento, todas as
// outras threads que quiserem usar aquilo vão ter que esperar na fila e esperar que aquele primeiro termine
// e depois eles vão poder executar.

// A gente vai ter um tipo Mutex, uma trava de exclusão mútua, então a gente vai precisar de um valor desse tipo.
// E a gente vai ter dois métodos: Lock que tranca e Unlock que destranca.
