package main

import "os"

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}

// Qual foi a diferença entre log.Panicln e panic? Que um me dá um log e o outro não me dá log nenhum somente o panic.
// A recomendação é sempre utilizar log e você opta de acordo com a severidade do seu programa.
