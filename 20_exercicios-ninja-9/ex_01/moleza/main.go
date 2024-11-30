package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		fmt.Println("Olá, eu sou a goroutine 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Olá, eu sou a goroutine 2")
		wg.Done()
	}()

	wg.Wait()
}
