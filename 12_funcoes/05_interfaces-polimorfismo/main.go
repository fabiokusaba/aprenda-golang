package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario          float32
}

type arquiteto struct {
	pessoa
	tipoconstrucao   string
	tamanhodaloucura string
}

func (x dentista) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e eu já arranquei", x.dentesarrancados, "dentes, e ouve só: Bom dia!")
}

func (x arquiteto) oibomdia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()

	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.(dentista).salario)
	case arquiteto:
		fmt.Println("Eu construo:", g.(arquiteto).tipoconstrucao)
	}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Mister Dente",
			sobrenome: "da Silva",
			idade:     50,
		},
		dentesarrancados: 8973,
		salario:          98736.06,
	}

	mrpredio := arquiteto{
		pessoa: pessoa{
			nome:      "Mister Prédio",
			sobrenome: "Sobrenome",
			idade:     51,
		},
		tipoconstrucao:   "Hospícios",
		tamanhodaloucura: "Demais",
	}

	// mrdente.oibomdia()
	// mrpredio.oibomdia()

	// Funciona igual, ao invés de eu ter que chamar o metodo específico de cada tipo eu consigo ter uma interface e a
	// interface usa valores de tipos diferentes, mas chama a mesma coisa dentro de cada valor.
	serhumano(mrdente)
	serhumano(mrpredio)
}

// Em Go valores podem ter mais que um tipo.
// Uma interface permite que um valor tenha mais que um tipo.
// Uma interface é um conjunto de metodos, por exemplo você diz que se o bixo late, rola na grama e come ração ele é um
// cachorro, então você tem uma interface cachorro com três metodos ali. Os tipos que você criar e que tiver aqueles três
// metodos vai fazer parte daquela interface, isso significa que ela vai ganhar o tipo da interface.
// Polimorfismo significa que eu vou poder através de uma mesma ação eu vou poder trabalhar em coisas de tipos distintos
// então ao invés de eu ter uma função pro cachorro e uma função pro gato eu vou poder ter uma mesma função que recebe os
// dois tipos através da interface do qual os dois fazem parte.
