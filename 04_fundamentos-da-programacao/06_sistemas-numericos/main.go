package main

import "fmt"

func main() {
	a := 100
	fmt.Printf("%d\t%b\t%#x", a, a, a)
}

// O nosso sistema numérico convencional vai de 0 a 9 e como é que ele funciona? A gente tem 9 símbolos a gente vai subindo
// 1, 2, 3, 4, quando a gente chega no último símbolo a gente sobe um valor na próxima casa.
// Então, a gente tem as unidades, dezenas, centenas, milhares, centenas de milhares, milhões, dezenas de milhões.
// Os sistemas binário e hexadecimal funcionam da mesma maneira só que com uma quantidade diferente de símbolos, por exemplo
// em binário a gente tem dois símbolos e hexadecimal são dezesseis.
// Base 10: decimal, 0 - 9
// Base 2: binário, 0 - 1
// Base 16: hexadecimal, 0 - f
