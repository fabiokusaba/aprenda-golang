package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "20julho1980"
	senhaerrada := "20julho1990"

	// Gerando um hash a partir da senha
	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	// Verificando o erro
	if err != nil {
		fmt.Println(err)
	}
	// Hash
	fmt.Println(string(sb))

	// Comparando senhas
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))
}

// bcrypt é uma função que faz um hashing de senhas.
// hashing - um hash vai pegar qualquer quantidade de dados que você der para a função vai utilizar uma fórmula
// matemática e vai te retornar um número e o único jeito de você obter aquele número basicamente/teoricamente é
// utilizando aquele mesmo arquivo.
