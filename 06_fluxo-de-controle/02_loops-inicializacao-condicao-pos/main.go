package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		fmt.Println(x)
	}
}

// Inicialização -> antes de começar a rodas as repetições do programa a gente inicializa com algum statement que a gente quiser.
// Condição -> é uma coisa que tem que ser verdade, ou seja, a gente vai ter um valor booleano que vai ser true para que o programa
// continue repetindo.
// Pós -> ao final de cada execução a gente tem mais uma ação que no caso é o pós
// Ponto e vírgula -> em Go ao final de todo statement existe um ponto e vírgula, você não precisa colocar ele no seu código mas o
// compilador vai colocar ele pra você, ele vai determinar o final de cada instrução.
// Para quem está vindo de outras linguagens não existe while em Go, nós temos somente for.
