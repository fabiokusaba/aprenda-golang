package main

import "fmt"

func main() {
	y := 24
	x := y << 2
	// y := x << 1
	z := y >> 2

	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", z)
}

// Deslocamento de bits é quando deslocamos dígitos binários da esquerda para direita ou da direita para esquerda.
// E o que isso quer dizer? Por exemplo:

// 0110000
// 0011000
// 0001100
// 0000110
// 0000011

// O operador de deslocamento de bits é esse aqui ">>" para um lado e "<<" para o outro lado.
