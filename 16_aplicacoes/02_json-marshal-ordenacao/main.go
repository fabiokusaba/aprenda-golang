package main

import (
	"encoding/json"
	"fmt"
)

// Para os campos de um struct serem importados para json eles tem que inicializar com letra maiúscula, em Go
// quando temos qualquer coisa iniciando com a letra maiúscula ela é exportada da nossa função, ou seja, eu posso
// ter outro package que vai importar o nosso package e quando ocorre essa importação tudo o que está em letra
// maiúscula é visível, pode ser importado, tudo o que está em letra minúscula não é.
type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissao     string
	Contabancaria float64
}

func main() {
	jamesbond := pessoa{
		Nome:          "James",
		Sobrenome:     "Bond",
		Idade:         40,
		Profissao:     "Agente Secreto",
		Contabancaria: 1000000.50,
	}

	darthvader := pessoa{"Anakin", "Skywalker", 50, "Vilão Intergaláctico",
		50000000000000.83}

	// Transformando objetos para json
	// Nós temos a função json.Marshal que me dá dois valores: slice of bytes e error
	j, err := json.Marshal(jamesbond)
	// Tratando erros
	if err != nil {
		fmt.Println(err)
	}
	d, err := json.Marshal(darthvader)
	// Tratando erros
	if err != nil {
		fmt.Println(err)
	}

	// Demonstrando os resultados
	fmt.Println(string(j))
	fmt.Println(string(d))
}
