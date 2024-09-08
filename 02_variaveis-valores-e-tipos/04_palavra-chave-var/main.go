package main

import "fmt"

var y = 10

func main() {
	z := 20
	qualquercoisa(z)
}

func qualquercoisa(x int) {
	fmt.Println(y)
	fmt.Println(x)
}

// A variável declarada em um codeblock é undefined em outro.
// Para você declarar uma variável com package level scope você precisa do var e você precisa declarar fora de um codeblock, como
// eu não posso usar a marmota fora de um codeblock eu sou obrigado a usar o var.
// O var funciona em qualquer lugar, ou seja, declarado em package level scope eu posso usar onde eu quiser.
// Prestar atenção em: chaves, parênteses e colchetes.
// Em resumo o var é utilizado fora do codeblock a gente pode utilizá-lo em qualquer situação em que a gente utilizaria o short
// declarator, mas sempre vamos utilizar a marmota quando possível e a gente vai utilizar o var somente quando a gente precisar de
// package level scope.
