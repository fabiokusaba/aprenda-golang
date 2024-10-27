package main

import "fmt"

// Declaração de uma struct.
type cliente struct {
	// E quando declaramos uma struct temos que declarar quais os tipos de dados diferentes que a gente quer que tenha
	// aqui dentro.
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	// Criando valores para aquele tipo.
	cliente1 := cliente{
		nome:      "João",
		sobrenome: "da Silva",
		fumante:   false,
	}

	cliente2 := cliente{"Joana", "Pereira", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}

// Struct é um tipo de dados cujo o nome vem da palavra inglesa estrutura (structure) e o struct nos permite armazenar
// dados com tipos diferentes, então eu posso ter dentro de um struct um string, integer, float, bool, etc.
