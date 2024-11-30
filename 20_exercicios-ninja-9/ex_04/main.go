package main

import (
	"fmt"
	"runtime"
	"sync"
)

var contador int
var wg sync.WaitGroup
var mu sync.Mutex

const quantidadeGoroutines = 5

func main() {
	contador = 0
	criarGoroutines(quantidadeGoroutines)
	wg.Wait()
	fmt.Println(contador)
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}
