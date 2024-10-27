package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	pessoas := make(map[string]pessoa)

	pessoas["da Silva"] = pessoa{
		nome:      "João",
		sobrenome: "da Silva",
		sabores:   []string{"Morango, Chocolate, Flocos"},
	}

	pessoas["Pereira"] = pessoa{"Maria", "Pereira", []string{"Pistache", "Napolitano"}}

	for _, v := range pessoas {
		fmt.Println("Meu nome é", v.nome, "e meus sorvetes favoritos são:")

		for _, s := range v.sabores {
			fmt.Println("-", s)
		}
	}

	//pessoas2 := map[string]pessoa{
	//	"da Silva": pessoa{
	//		nome:      "João",
	//		sobrenome: "da Silva",
	//		sabores:   []string{"Morango, Chocolate, Flocos"},
	//	},
	//	"Pereira": pessoa{"Maria", "Pereira", []string{"Pistache", "Napolitano"}},
	//}

	//fmt.Println(pessoas2)

}
