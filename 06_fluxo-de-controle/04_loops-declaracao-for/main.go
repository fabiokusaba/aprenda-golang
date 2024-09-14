package main

import "fmt"

func main() {
	x := 0
	for x < 10 {
		fmt.Println("chis é menor que déis")
		x++
	}

	for {
		if x < 10 {
			x++
			fmt.Println("chis é menor que déis")
		} else {
			fmt.Println("chis é maior que déis mano tô fora!")
			break
		}
	}
	fmt.Println("o loop está brekado!")
}

// O que é um while? While é um loop que enquanto uma condição for verdadeira ele faz determinada coisa e sabemos que em Go não
// temos uma estrutura loop assim, em Go nós temos o for.
// O for em Go pode ser usado com as três instruções de inicialização, condição e pós ou ele pode ser usado somente com uma condição.
// Qual a utilidade de um processo que roda pra sempre? Por exemplo se você tem um servidor de HTTP para servir páginas ele fica ali
// o dia inteiro dizendo "alguém está pedindo página?", ele fica rodando a vida inteira, e cada vez que alguém pede alguma coisa ele
// está lá "alguém está pedindo página? Sim, ele está pedindo página. Então, vou dar página para essa pessoa. Está entregue? Ok."
// A condição não é necessária, ela pode não existir, ou seja, se eu quero um loop infinito que vai ficar rodando pra sempre como num
// servidor HTTP eu posso declarar o for sem a condição.
// O break ele quebra qualquer loop, ele simplesmente diz "parou de repetir isso daqui vamos sair fora e seguir o programa".
// O statement for especifica a execução repetidamente de um bloco de código. Ou seja, a repetição pode ser controlada por uma única
// condição que é aquele "while", por uma "for clause" que é aquilo que vimos com os três itens: inicialização, condição e pós, ou
// por um "range clause".
