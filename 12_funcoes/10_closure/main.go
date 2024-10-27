package main

import "fmt"

func main() {
	// A nossa variável "a" vai ser igual ao resultado da nossa função "i" e o resultado dessa função é uma outra função
	// então "a" vai se tornar essa função.
	a := i()

	// Executando a função "a" e posso fazer isso várias vezes, ou seja, cada vez que ele é executado ele incrementa 1
	// e retorna o valor.
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

// Closure é quando a gente captura um scope e sequestra ele e usa ele de refém pra fazer todas as nossas ordens.
// É cercar ou capturar um scope para que possamos utilizá-lo em outro contexto, ou seja, a gente captura um contexto.
func i() func() int {
	// Aqui nós temos um closure e com o closure a gente captura um scope pra utilizar o scope somente quando a gente
	// quiser, ou seja, quando essa função é criada ela utiliza uma variável que está fora do scope dela, que está no
	// scope da função subjacente, o "x" não está na função ele está fora ele está em outro scope, eu consigo capturar
	// esse "x" e utilizar ele aqui dentro apesar dele estar fora e esse "x" vai ser único cada vez que eu fizer um
	// closure e capturar um scope de fora da função eu vou ter uma cópia diferente desse scope.
	x := 0
	return func() int {
		x++
		return x
	}
}
