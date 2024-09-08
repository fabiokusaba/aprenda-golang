package main

import "fmt"

type hotdog int

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", b)

	// Conversão
	x = int(b)
	fmt.Printf("%v\n", x)
}

// Conversão de tipos é o que o nome quer dizer, é você converter um tipo em outro tipo as vezes isso é possível e as vezes não é.
// Na aula passada a gente tinha criado um tipo chamado "hotdog" o seu tipo subjacente é um int e a gente viu que quando usamos os
// dois de maneira igual o resultado deles é igual, então por que não seria compatível? Ele não é compatível porque os tipos são
// diferentes apesar do valor por trás dos panos ser o mesmo, mas eu posso utilizar esse valor.
// Ou seja, para eu converter um valor para outro tipo eu simplesmente coloco o tipo que eu quero e o valor entre parênteses logo em
// seguida T(x).
