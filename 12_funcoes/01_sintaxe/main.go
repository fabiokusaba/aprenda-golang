package main

import "fmt"

func main() {
	basica()

	argumento("tarde")

	valor := soma(10, 10)
	fmt.Println(valor)

	total, quantosElementos, saudacao := soma2("tarde", 10, 10, 1, 2, 3, 5)
	fmt.Println(total, quantosElementos, saudacao)

}

// As funções elas servem pra abstrair funcionalidades, descomplicar as coisas para que tarefas específicas fiquem mais
// fáceis.
// Por esse motivo uma das outras razões que se utilizam funções é para reutilização de código.
// Toda função segue essa estrutura:
// `func (receiver) identifier(parameters) (returns) { code }`
// A diferença entre parâmetros e argumentos: as funções são definidas com parâmetros e são chamadas com argumentos.
// Tudo em Go é "pass by value", ou seja, é passado o valor para ser executado pela função.

// Função básica
func basica() {
	fmt.Println("Oi, bom dia!")
}

// Função que aceita um parâmetro
func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi, bom dia!")
	} else if s == "tarde" {
		fmt.Println("Oi, boa tarde!")
	} else {
		fmt.Println("Oi, boa noite!")
	}
}

// Função com retorno
// Quando os parâmetros são do mesmo tipo posso escrevê-los da seguinte forma: "x int, y int" ou "x, y int".
func soma(x, y int) int {
	return x + y
}

// Função com múltiplos retornos
// Vou ter um parâmetro variádico de ints (slice of ints) e vou retornar dois ints.
// Quando você tem uma função com parâmetros variádicos o elemento variádico do parâmetro tem que ser o último, você
// pode ter quantos elementos que você quiser antes do parâmetro variádico, mas o variádico tem que ser o último.
func soma2(s string, x ...int) (int, int, string) {
	saudacao := ""

	if s == "manhã" {
		saudacao = "Oi, bom dia!"
	} else if s == "tarde" {
		saudacao = "Oi, boa tarde!"
	} else {
		saudacao = "Oi, boa noite!"
	}

	soma := 0

	for _, v := range x {
		soma += v
	}

	return soma, len(x), saudacao
}
