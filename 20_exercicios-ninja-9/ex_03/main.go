package main

import (
	"fmt"
	"runtime"
	"sync"
)

var contador int
var wg sync.WaitGroup

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
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
}
