package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

func main() {
	mauricio := pessoa{"Maurício", 30}
	mauricio.oibomdia()
}

// Um metodo é uma função anexada a um tipo e quando anexamos uma função a um tipo ela se torna um metodo daquele tipo.
// O receiver vai dizer o que a função recebe é tipo um parâmetro, mas é um pouquinho diferente porque o receiver é como
// se fosse exclusivo, ele vai dizer que a função só pode receber o valor desse tipo e ela só está anexada aquele tipo
// enquanto que o argumento ele recebe valores, o receiver é mais do que valores porque ele diz o contexto da função.
// O metodo adiciona funcionalidade especificamente a um tipo.
