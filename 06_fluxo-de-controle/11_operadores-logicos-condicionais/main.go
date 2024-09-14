package main

import "fmt"

func main() {
	x := 4

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("é múltiplo de dois e tambem de treis")
	}

	if x%2 == 0 || x%3 == 0 {
		fmt.Println("é múltiplo de dois ou de treis")
	}
}

// Operadores lógicos condicionais são: && que quer dizer E, || que quer dizer OU.
// Como vimos isso daqui é uma expressão booleana, ou seja, o resultado disso daqui vai ser true se o resultado de alguma das
// expressões for verdadeiro. Ou seja, eu tenho um boolean que é true ou false, tenho outro boolean que é true ou false aquele OU
// significa que se esse ou esse for verdade dá um resultado verdadeiro. Se esse ou esse forem mentira os dois dá um resultado
// negativo.
// Esse operador E só vai dar um true se as duas expressões forem true.
