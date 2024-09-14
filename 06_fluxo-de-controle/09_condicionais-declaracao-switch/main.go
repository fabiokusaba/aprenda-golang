package main

import "fmt"

func main() {
	x := 10

	switch {
	case x < 5:
		fmt.Println("chis é menor que cinco")
	case x == 5:
		fmt.Println("chis é igual a cinco")
	case x > 5:
		fmt.Println("chis é maior que cinco")
	}

	quemtanoescritoriohoje := "zezinho"

	switch quemtanoescritoriohoje {
	case "zezinho", "maria":
		fmt.Println("hoje quem tá na firma é o time 1")
	case "joana", "pedrinho":
		fmt.Println("hoje quem tá na firma é o time 2")
	default:
		fmt.Println("tá vazio. ou a balada tava ótima, ou é feriado")
	}
}

// O switch é como se fosse um if com vários casos diferentes, mas o foco dele é realmente ter vários casos diferentes não é
// simplesmente ver se uma expressão é verdadeira ou não.
// Ele avalia expressões e escolhe qual está correta e executa de cima para baixo.
// Uma coisa interessante da linguagem Go comparada com outras linguagens é que não existe fall-through por padrão, o que é
// fall-through? Significa que se um case está certo o próximo também vai estar, por exemplo quando um caso com fallthrough é
// executado eu pulo essa comparação e executo direto o próximo caso.
// Se nada daquilo acontecer, ou seja, se ele não entrar em nenhum case o default será executado. Faço as comparações do switch
// statement com o case statement caso eu não encontre nada eu vou executar o default.
// Cases compostos -> ou seja, a expressão para ele fazer a comparação aqui não necessariamente precisa ser uma você pode ter
// várias.
