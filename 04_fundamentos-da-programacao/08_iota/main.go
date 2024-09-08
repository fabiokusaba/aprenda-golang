package main

import "fmt"

const (
	a = iota * 10
	_
	c
	x
	_
	z
)

func main() {
	fmt.Println(a, c, x, z)
}

// Iota -> dentro de uma declaração de uma constante o identificador pré-declarado iota representa sucessivos números inteiros
// que não necessariamente estão tipados ainda, ou seja, a mesma coisa que fizemos no exemplo anterior vou ter um número ali vou
// escrevê-lo como um integer mas ele só vai ser tipado no momento do uso, ou seja, eu posso usar ele como float.
// E para que isso serve? Ele serve para situações em que você não necessariamente está ligando para qual é o valor de uma constante
// você só quer que ela seja diferente das outras.
// É possível descartar valores com underscore, underline aquele mesmo caracter que já vimos.
