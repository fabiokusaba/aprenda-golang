package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	x := 10
	fmt.Printf("%T\n", x)
	fmt.Printf("%T", b)
}

// Em Go tipos são extremamente importantes.
// Em Go tipos são imutáveis se eu declarar uma variável como int ela vai ser int até ela morrer.
// A gente criou um tipo customizado, definido pelo usuário, chamado "hotdog" e esse tipo ele tem um tipo subjacente e o que quer
// dizer subjacente? É algo que vem por trás como base, ou seja, o tipo base do tipo "hotdog" é um int.
// Quando criamos os nossos próprios tipos podemos fazer muitas coisas com eles que a gente não pode fazer com os tipos que já vem
// na linguagem.
