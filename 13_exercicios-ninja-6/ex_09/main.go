package main

import "fmt"

func main() {
	executarSaudacao(saudacao, "Maria")
}

func saudacao(nome string) {
	fmt.Println("Olá,", nome, "prazer em te conhecer.")
}

func executarSaudacao(f func(string), nome string) {
	f(nome)
}
