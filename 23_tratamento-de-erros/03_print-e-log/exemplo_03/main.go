package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")
}

// Eu peguei esse ponteiro para um arquivo aqui 'f' e coloquei no SetOutput, que toma um io.Writer, então uma interface
// Writer e qualquer Writer você pode usar, aqui estamos usando o arquivo mas pode ser qualquer uma das um milhão de
// coisas que usam a interface Writer.
// Aqui estamos dizendo que quando ocorrer alguma coisa e o package log tiver que salvar alguma coisa salva aqui, como a
// gente está salvando para um ponteiro para um arquivo então ele vai pegar o conteúdo e escrever no arquivo, poderia
// aqui no lugar colocar o stdout que é a tela.
