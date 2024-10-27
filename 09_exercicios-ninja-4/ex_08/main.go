package main

import "fmt"

func main() {
	pessoas := map[string][]string{
		"silva_maria":      []string{"Cozinhar", "Assistir séries e filmes"},
		"roberto_jorge":    []string{"Assistir futebol", "Happy hour com os amigos"},
		"felizberto_roger": []string{"Caminhar", "Ler bons livros"},
	}

	for k, v := range pessoas {
		fmt.Println(k)

		for i, hobby := range v {
			fmt.Println("\t", i, hobby)
		}
	}

}
