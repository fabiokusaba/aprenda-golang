package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

// Documentação JSON
// JSON (Javascript Object Notation) - ou seja, é uma notação de objetos que foi inicialmente utilizada em Javascript
// e que hoje em dia se usa para quase tudo.
// O pacote json implementa a codificação e decodificação de JSON conforme definido no RFC 4627.
