package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")

}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// Tudo isso daqui é pra basicamente você entender que quando a gente faz um panic a gente tem essa história de ter
// várias funções, tipo a função 1 chamou a função 2 que chamou a 3 que chamou a 4 e assim por diante, então o panic vai
// encerrar elas em ordem reversa, vai rodar todos os defer até encontrar um recover, quando encontrar um recover aí a
// gente corta esse fluxo do panic e volta para uma execução normal.
