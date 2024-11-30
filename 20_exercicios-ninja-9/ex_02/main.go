package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println("Olá, me chamo", p.nome, "e tenho", p.idade, "anos.")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {
	p := pessoa{
		nome:  "John Doe",
		idade: 25,
	}

	// Se x for endereçável e o conjunto de metodos de ponteiro para x contiver o metodo, então x.m() é um atalho
	// para (&x).m()
	p.falar()    // É um shortcut pra (&p).falar()
	(&p).falar() // É a maneira "mais correta."

	// Ponteiro para pessoa tem o metodo falar() que faz parte da interface humanos e por conta disso eu posso usar
	// o metodo dizerAlgumaCoisa que recebe como parâmetro humanos, enquanto que se eu passar somente uma pessoa
	// ela não possui o metodo falar(), não faz parte da interface humanos e por conta disso não posso utilizar a
	// função dizerAlgumaCoisa que toma um humanos como parâmetro.
	dizerAlgumaCoisa(&p) // Funciona, estou passando um ponteiro para pessoa.
	//dizerAlgumaCoisa(p) // Não funciona, estou passando apenas pessoa.
}
