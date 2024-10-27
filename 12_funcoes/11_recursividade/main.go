package main

import "fmt"

func main() {
	fmt.Println(fatorial(4))
	fmt.Println(fatorialLoop(4))

}

// A aplicação mais comum de recursividade é em matemática e ciência da computação onde a definição de uma função é a
// aplicação da própria função nela mesma.
func fatorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * fatorial(n-1)
}

func fatorialLoop(n int) int {
	total := 1
	if n == 0 || n == 1 {
		return 1
	}
	for n > 1 {
		total *= n
		n--
	}
	return total
}
