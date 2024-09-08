package main

import "fmt"

// Declaração
var x int
var a int
var b float64
var c string
var d bool

func main() {
	// Inicialização e atribuição
	x = 10
	fmt.Printf("%v, %T\n", x, x)

	// Atribuição
	x = 20
	fmt.Printf("%v, %T\n\n", x, x)

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}

// Declaração vs inicialização vs atribuição de valor a variáveis -> vamos utilizar o exemplo de uma caixa postal. Eu vou no
// correio e digo ao funcionário que quero uma caixa postal, o funcionário vai me entregar vários documentos para assinar, vou
// fazer o pagamento de uma taxa e ao final vou ter a minha caixa postal, ou seja, vou ter um endereço onde vou poder colocar
// coisas dentro, isso chama-se declaração. A inicialização é o primeiro valor que eu coloco dentro de uma variável, ou seja,
// comprei uma caixa postal a primeira carta que eu receber, pacote ou seja lá o que for, a primeira coisa que for colocada nessa
// caixa postal isso vai se chamar inicialização. Então, declaração é quando eu compro um endereço de caixa postal, inicialização
// é quando eu recebo a primeira correspondência nessa caixa postal e atribuição é o valor que é colocado depois, por exemplo eu
// tenho na minha caixa postal um x resolvo pegar esse x e jogar fora e colocar um y isso é uma atribuição.
// O operador curto de declaração tem esse nome porque ele faz tudo ao mesmo tempo.
// O valor zero é o valor que se encontra presente numa variável antes dela ser inicializada pelo usuário, ou seja, quando você
// compra uma caixa postal ela vem vazia e esse vazio é um zero e esse zero é o que a gente quer ver. Se a sua variável for uma
// string o seu valor zero vai ser um, se a sua variável for um integer o seu valor zero vai ser outro, se a sua variável for um
// float vai ser outro, se for um bool vai ser outro, ou seja, cada tipo da linguagem vai ter um valor zero diferente.
// Os valores zeros são: int 0, float 0.0, bool false, string "". Pointers, functions, interfaces, slices, channels, maps nil.
// Sempre que possível utilize a marmota, operador curto de declaração.
// E quando a gente precisar de package level scope é a situação em que a gente não vai utilizar a marmota e a gente vai utilizar o
// var.
// Sempre que você declarar uma variável e não inicializá-la ela vai conter por padrão o valor zero.
