package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type informacoes struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Trabalho"`
	Contabancaria float64 `json:"Contabancaria"`
}

func main() {
	sb := []byte(
		`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":1000000.50}`)

	var jamesbond informacoes
	err := json.Unmarshal(sb, &jamesbond)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(jamesbond)
	fmt.Println(jamesbond.Profissao)

	// Utilizando a função Encoder
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesbond)
}

// Unmarshal analisa a codificação em JSON e salva o resultado no valor apontado pela variável v, ou seja, isso
// aqui tem que ser um ponteiro, se v for nada ou se ele não for um ponteiro ele retorna um erro.
// Para fazer um unmarshal dentro de uma struct eu preciso que os campos da minha struct sejam equivalentes aos
// campos que eu vou ter no JSON.
// As tags, por exemplo `json:"Nome"`, servem para quando você recebe uma informação em JSON e ela tem um nome no
// JSON e no seu código tem outro nome, por exemplo suponhamos que no JSON tenha o nome "Trabalho" e no meu código
// tenha o nome "Profissao".
// A diferença é que quando eu executo o Marshal o resultado vai para uma variável e depois eu faço o que eu quiser
// com essa variável, por exemplo um print. O Encoder quando eu executo o Encode ele vai direto para a interface que
// eu coloquei aqui, então não preciso salvar numa variável intermediária pra depois mandar pra lá, ele vai direto.
