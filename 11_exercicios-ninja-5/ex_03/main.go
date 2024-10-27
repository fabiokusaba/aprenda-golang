package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	ford := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "branco",
		},
		tracaoNasQuatro: true,
	}

	toyota := sedan{
		veiculo: veiculo{
			portas: 5,
			cor:    "prata",
		},
		modeloLuxo: true,
	}

	fmt.Println(ford)
	fmt.Println(toyota)

	fmt.Println(ford.cor)
	fmt.Println(toyota.modeloLuxo)
}
