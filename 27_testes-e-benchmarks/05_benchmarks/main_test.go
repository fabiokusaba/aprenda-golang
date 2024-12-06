package main

import (
	funcionalidades "aprenda-golang/27_testes-e-benchmarks/02_testes-em-tabela"
	"testing"
)

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funcionalidades.Soma(3, 2, 1)
	}
}

func BenchmarkMultiplica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funcionalidades.Multiplica(2, 1)
	}
}

// São medidas de performance ou velocidade pra ver qual código funciona melhor porque você tem uma função que pode ser
// escrita de 20 maneiras diferentes, então qual é a melhor maneira? Qual a maneira com a melhor performance? Pra obter
// a resposta pra essa pergunta a gente utiliza benchmarks.
// Na prática vai no arquivo _test.go
// BET: em inglês quer dizer aposta, então o que eles dizem é assim "eu posso apostar nesse código?" E o que isso quer
// dizer? Se esse código tem BET (Benchmarks, Examples, Test), se o meu código tiver isso então esse é um código no qual
// posso apostar.
