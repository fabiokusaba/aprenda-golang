package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened", err)
	}
}

// O Println funciona do jeito que já sabemos e o resultado sai na tela.
