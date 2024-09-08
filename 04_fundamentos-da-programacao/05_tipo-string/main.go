package main

import "fmt"

func main() {
	s := "ascii éâ u999"
	// sb := []byte(s)

	for _, v := range s {
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}

// Strings são sequências de bytes.
// Strings são imutáveis, então em Go eu não consigo pegar uma string e alterar ela, no caso eu teria que criar outra string
// contendo a string anterior com a alteração.
// Uma string é um "slice of bytes".
// Então, coisas que a gente tem que saber se a gente faz um range numa string ele não vai me dar byte por byte ele vai me dar
// caracter por caracter, se a gente fizer com o for loop ele vai me dar por byte.
