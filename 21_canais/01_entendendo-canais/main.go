package main

import "fmt"

func main() {
	// Criando um canal
	//canal := make(chan int)
	// Com buffer
	canal := make(chan int, 1)
	// Colocando um valor no buffer do canal
	canal <- 42
	// deadlock!
	// canal <- 43
	// Transformando em uma goroutine
	// go func() {
	// 	canal <- 42
	// }()
	// Exibindo o valor
	fmt.Println(<-canal)
}

// Canais são uma das características mais interessantes da linguagem Go e são a maneira da gente transmitir
// dados entre goroutines, lembrando que a função main() também é uma goroutine.
// Então, vamos pegar várias go funcs e a gente vai transmitir informações entre elas e dessa maneira a gente
// consegue fazer sincronização e código concorrente do jeito certo.
// Canais bloqueiam: eles são como corredores em uma corrida de revezamento, eles tem que passar o bastão de
// maneira sincronizada, se o corredor tentar passar o bastão para o próximo mas o próximo não estiver lá dá
// problema, ou se um corredor ficar esperando para receber o bastão mas ninguém entregar dá problema.
// Então, o que aconteceu no trecho de código acima é que eu coloquei o 42 no canal só que para eu largar o 42
// lá alguém tem que pegar tipo uma corrida de revezamento onde um corredor entrega o bastão para outro corredor
// ele não pode largar o bastão no chão e ir fazer outra coisa.
// O nosso Println() nunca chegou a rodar porque a execução ficou presa no 42, então a conclusão que a gente tira
// que colocar informação no canal e retirar informação no canal tem que ser feita de maneira concorrente, ou seja
// eu preciso de uma goroutine entregando informação pra outra goroutine. Eu não posso numa mesma goroutine colocar
// e tirar informação simultâneamente, então preciso de duas.
// O buffer é o seguinte: posso pegar um canal e dizer pra ele que esse canal não preciso que alguém retire
// informação ao mesmo tempo que alguém põe informação, eu posso ter o buffer de um e isso quer dizer que eu posso
// deixar um valor ali dentro e a pessoa pode tirar depois para tirar um pouco dessa sincronia.
// Os buffers também bloqueiam porque eles possuem um valor finito, então
