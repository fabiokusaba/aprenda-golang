package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) // zero value
	x = true
	fmt.Println(x) // valor atribuído

	x = (10 < 100)
	y := (10 == 100)
	z := (10 > 100)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

// O tipo bool é um tipo binário, ou seja, que só pode conter dois valores: true e false.
// Sempre que você ver operadores relacionais (==, <=, >=, !=, <, >), o resultado da expressão será um valor booleano.
// Booleans são fundamentais nas tomadas de decisões em lógica condicional, declarações switch, declarações if, fluxo de controle, etc.
