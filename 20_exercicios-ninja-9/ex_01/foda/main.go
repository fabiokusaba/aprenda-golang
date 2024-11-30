package main

import (
	"fmt"
	"sync"
)

// Declarando WaitGroup em package level scope
var wg sync.WaitGroup

func main() {
	novasgoroutines(100)
	wg.Wait()
}

// Vou rodar x vezes esse codeblock, esse codeblock vai criar uma variável nova aqui dentro que vai servir como closure
// e vai ser o número da iteração que está aqui, vou passar esse número como argumento para a minha go func e a minha
// go func vai ser um print na tela
func novasgoroutines(i int) {
	wg.Add(i)
	for j := 1; j <= i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Eu sou a goroutine número:", i)
			wg.Done()
		}(x)
	}
}
