package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Panicln(err)
	}
}

// Quando você chama Panicln é como se você tivesse chamado Println com uma mensagem e um panic().
// E Fatalln é equivalente a Println seguido por uma chamada de os.Exit(1).
// A função panic para a execução normal da goroutine atual, então quando uma funçã F chama panic a execução da função F
// é interrompida imediatamente.
// Ou seja, com o panic funções com defer rodam, então está aí uma diferença bem significativa do Fatalln, então quando
// eu chamo panic eu paro a execução onde eu estou rodo o que tem de defer e eu fecho a função atual e retorno.
// Então, se eu tinha uma função main dentro da main chamei a função qualquercoisa() e essa função deu um panic o que
// vai acontecer? Eu vou parar a execução da função qualquercoisa() vou rodar o que tiver de defer na função qlqrcosia()
// vou retornar essa função, a minha execução vai cair de volta na função main, a função main vai entrar em panic e o
// que quer dizer? Ela para a sua execução, ela roda o que tiver de defer e ela retorna também, no caso quando a gente
// retorna da função main o programa fecha.
// Essa sequência de terminação se chama panicking, entrar em pânico, e pode ser controlada pela função recover().
