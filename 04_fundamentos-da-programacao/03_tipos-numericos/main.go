package main

import "fmt"

var x = 10
var y = 10.0

func main() {
	a := "e"
	b := "é"
	c := "u999"
	fmt.Printf("%v %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v, %v, %v\n\n", d, e, f)

	// x := 10
	// y := 10.0
	x = 20
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
}

// Existem dois tipos numéricos principais que são: int e float. Int são números inteiros, exemplo 10, 20, 30. Float são
// números com fração, ou seja, 10.5, 1/3, 89.7134, PI.
// A diferença entre uint (unsigned) e int (signed) é o sinal de menos, uint8 (unsigned int) eu tenho 8 bits é um byte de
// 8 bits que pode guardar um valor entre 0 e 255, o int8 (signed int) ele tira um desses bits pra dizer se é positivo ou
// negativo, ou seja, eu tenho 7 bits numéricos e 1 bit que é o sinal, 7 bits como a gente já tinha visto é um número de 0
// até 128, então a gente tem 128 até -1 e 0 até 127.
// float são números com vírgula.
// complex são números complexos.
// byte é um alias para uint8.
// rune é um alias para int32.
// Esses números quando você declara são específicos ao seu computador, por exemplo se eu disser: "var x int = 10", esse int
// vai ser um int32 ou um int64? Isso depende do seu computador, vai ser do tamanho da palavra do seu computador, palavra/word
// determina quantos bits o seu processador interpreta de uma vez só. Ou seja, o tamanho do número que você vai passar para o
// seu processador vai ser do tamanho da palavra que é quantos bits o processador aguenta, ou seja, quando eu faço "var x int = 10"
// esse int vai ser int32 num computador de 32 bits, vai ser um int64 num computador de 64 bits.
// Todos os tipos numéricos são distintos, exceto: byte e rune. Ou seja, 8 bits é um byte então eu tenho aqui um apelido para uint8
// mesma coisa o rune, strings são feitas de runes cada rune é um caracter esse caracter pode ter 1, 2, 3 ou 4 bytes, então quando
// você quer mexer com caracteres UTF-8 ao invés de você utilizar int32 toda hora tem simplesmente um apelido que você pode colocar
// rune.
// Em Go tipos são únicos.
// Go é uma linguagem estática, ou seja, uma vez que você declara uma variável com tipo esse tipo não muda nunca.
// int e int32 não são a mesma coisa e para misturá-los é necessário conversão.
// Como regra geral quando você precisar de um int use somente int e deixe que o computador escolha pra você se é mais eficiente usar
// um ou usar o outro.
// Floating point são números racionais ou reais, ou seja, números com vírgula.
// Como regra geral use somente float64.
// Dá pra colocar número com vírgula em tipo int? Não, não consigo, ou seja, são tipos numéricos sim, mas não são necessariamente
// compatíveis a gente tem que fazer uma conversão pra poder utilizar um tipo como sendo outro.
