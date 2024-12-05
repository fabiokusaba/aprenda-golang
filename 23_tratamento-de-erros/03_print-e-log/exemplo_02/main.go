package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err happened", err)
	}
}

// O log.Println tem um timestamp, ou seja, quando dá um erro esse erro fica gravado com a data e o horário de quando
// ocorreu esse erro, uma coisa comum para qualquer tipo de log.
