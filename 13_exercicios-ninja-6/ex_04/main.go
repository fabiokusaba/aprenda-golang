package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) saudacao() {
	fmt.Println("Olá, me chamo", p.nome, p.sobrenome, "e tenho", p.idade, "anos.")
}

func main() {
	p1 := pessoa{
		nome:      "João",
		sobrenome: "da Silva",
		idade:     27,
	}
	p1.saudacao()
}
