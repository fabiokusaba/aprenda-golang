package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var contador int32
var wg sync.WaitGroup

const quantidadeGoroutines = 5

func main() {
	criarGoroutines(quantidadeGoroutines)
	wg.Wait()
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			// Incrementa o valor do contador
			atomic.AddInt32(&contador, 1)
			// Pega o valor do contador
			v := atomic.LoadInt32(&contador)
			// Entrega a thread
			runtime.Gosched()
			// Demonstre o valor
			fmt.Println(v)
			wg.Done()
		}()
	}
}
