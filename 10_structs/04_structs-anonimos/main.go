package main

import "fmt"

func main() {
	// Defini uma estrutura
	x := struct {
		nome  string
		idade int
	}{
		// Utilizei a estrutura
		nome:  "Mario",
		idade: 50,
	}

	// Imprimi a estrutura
	fmt.Println(x)
}

// Em um struct normal você declara o tipo e faz quantos objetos você quiser, quantos valores você quiser daquele tipo.
// O struct anônimo você não declara o tipo você simplesmente mostra uma estrutura e cria um valor que tem aquela
// estrutura e você não pode reutilizar várias vezes, é uma coisa descartável.
