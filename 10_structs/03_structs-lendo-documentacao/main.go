package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	// Aqui chamamos de um struct anônimo porque ele não tem o tipo aqui, ele está simplesmente apontando para o tipo
	// acima, mas ele não tem um nome específico.
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "Pizzaiola",
		salario: 10000,
	}

	// Declaração de forma concisa, sem precisar especificar o nome de cada campo.
	pessoa3 := pessoa{"Mauricio", 30}
	pessoa4 := profissional{pessoa{"Vanderlei", 50}, "Político", 10000000}

	// Acessando valores de um struct.
	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.titulo)
	fmt.Println(pessoa2.pessoa.nome)

	fmt.Println(pessoa3.nome)

	fmt.Println(pessoa4)

	// Caso não tenha conflito dos nomes dos campos serem iguais dizemos que o campo foi promovido, ou seja, ele foi de
	// um subcampo para o campo externo, então eu consigo acessar direto uma coisa que está em um struct embutido
	// chamando o struct inteiro.
	fmt.Println(pessoa2.idade)

}

// Tipos são a peça chave na linguagem Go, são as peças mais importantes.
// Struct é uma sequência de elementos com nome chamado campos cada um tem um nome e um tipo.
// Eles podem ser especificados explicitamente ou implicitamente.
// Nomes tem que ser únicos, não posso ter nome, nome, nome...
