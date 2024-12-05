package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}

// O Fatalln ele chama como está ali os.Exit(1).
// Exit faz com que o programa feche e dá um status code (código de status, um número), convencionalmente um código zero
// indica sucesso, e um código não zero significa erro, ou seja, o programa acaba imediatamente, funções com defer não
// rodam.
