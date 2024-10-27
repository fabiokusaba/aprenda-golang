package main

import "fmt"

func main() {
	x := retornaFuncao()
	fmt.Println(x())
}

func retornaFuncao() func() string {
	return func() string {
		return "Retornando uma função de uma função."
	}
}
