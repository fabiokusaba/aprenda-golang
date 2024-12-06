package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	y := multiplica(10, 10)
	fmt.Println(x)
	fmt.Println(y)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v // total = total + v
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v // total = total * v
	}
	return total
}

// Testes: a gente pega pedaços do nosso código e a gente, por exemplo, se eu tenho uma função que soma eu sei que se
// eu tiver dois valores 1 e 1 o resultado tem que ser 2, então eu faço testes pra conferir quando eu tento somar 1 + 1
// e o resultado vai ser 2, e ao longo do desenvolvimento do meu programa se eu alterar a minha função e em algum
// momento 1 + 1 não for 2 o meu teste vai falhar, vai dar negativo, então eu sei que alguma coisa ficou estranha e eu
// tenho que arrumar.
// Então, testes servem pra isso pra gente conferir se o nosso código está fazendo aquilo que a gente quer.
// TDD - Test Driven Development
// Testes devem ficar num arquivo cuja terminação seja _test.go
// Testes devem ficar na mesma package que o código a ser testado
// Testes devem ficar em funções com o nome "func TestNome(*testing,T)"
// Para rodar os testes a gente vai usar go test
// Para falhas a gente vai usar o t.Error() onde a maneira idiomática é algo do tipo "expected: x. got: y."
