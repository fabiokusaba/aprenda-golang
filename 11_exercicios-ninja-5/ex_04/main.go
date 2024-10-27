package main

import "fmt"

func main() {
	anonimo := struct {
		modeloMarca map[string]string
		cores       []string
	}{
		modeloMarca: map[string]string{
			"Corolla": "Toyota",
			"HB20":    "Hyundai",
			"Gol":     "Wolkgswagen",
		},
		cores: []string{"Branca", "Prata", "Preto"},
	}

	fmt.Println(anonimo)
}
