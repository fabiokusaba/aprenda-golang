package main

import "fmt"

func main() {
	// Um callback é passar uma função como argumento, por exemplo aqui nessa chamada de função um dos argumentos é um
	// slice of ints com vários números e o outro é a soma, a soma é uma função.
	t1 := somentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	t2 := somenteImpares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t1)
	fmt.Println(t2)

}

func soma(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}

func somentePares(f func(x ...int) int, y ...int) int {
	// Basicamente você executa uma funcionalidade.
	var pares []int
	for _, v := range y {
		if v%2 == 0 {
			pares = append(pares, v)
		}
	}
	// Quando essa funcionalidade acabar você chama uma função, essa função é como se fosse o ligue de volta, ou seja,
	// faz isso quando isso aqui acabar.
	total := f(pares...)
	return total
}

func somenteImpares(f func(x ...int) int, y ...int) int {
	var impares []int
	for _, v := range y {
		if v%2 != 0 {
			impares = append(impares, v)
		}
	}
	total := f(impares...)
	return total
}
