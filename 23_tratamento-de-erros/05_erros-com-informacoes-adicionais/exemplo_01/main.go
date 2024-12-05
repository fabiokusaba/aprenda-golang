package main

import (
	"errors"
	"log"
	"math"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

// Para que as nossas funções retornem erros customizados a gente pode utilizar: errors.New() e a gente pode usar um
// fmt.Errorf() que na verdade tem um errors.New() embutido.
// Em Go valores de erros não são especiais, eles são valores como qualquer outro, ou seja, você tem uma string, você
// tem um int, você tem um erro, então é um valor comum e você pode usar a linguagem inteira para tratar com esses erros
// do mesmo jeito que você gerenciaria qualquer outro valor. (Rob Pike)
