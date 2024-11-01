package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(p *pessoa) {
	(*p).nome = "Maria"
	p.sobrenome = "Luiza"
}

func main() {
	p1 := pessoa{
		nome:      "José",
		sobrenome: "da Silva",
		idade:     53,
	}

	fmt.Println(p1)
	mudeMe(&p1)
	fmt.Println(p1)
}
