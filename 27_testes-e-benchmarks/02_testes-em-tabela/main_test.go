package main

import "testing"

type test struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		test{data: []int{1, 2, 3}, answer: 6},
		test{[]int{10, 11, 12}, 33},
		test{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		z := Soma(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}

// Podemos escrever testes em série para testar variedades de situações.
// É a mesma lógica do teste simples só que ao invés de você testar uma execução da sua função você pega um monte de
// informação e executa quantas vezes você quiser com vários valores diferentes.
