package main

import "fmt"

func main() {
	// Canal bidirecional onde posso utilizar para qualquer coisa
	canal := make(chan int)

	// Transformando a função send() em uma go func
	go send(canal)

	receive(canal)
}

// Essas duas funções tem que rodar de maneira concorrente, ou seja, eu tenho que ter as duas rodando ao mesmo
// tempo, eu não posso rodar uma e depois a outra então uma das duas tem que ser uma go func.
// A minha função send() vai receber como parâmetro um canal onde vou colocar informação nele
func send(s chan<- int) {
	s <- 42
}

// Essa função receive() eu vou tirar valores do canal
func receive(r <-chan int) {
	fmt.Println("O valor recebido do canal foi:", <-r)
}

// Canais podem ser direcionais, ou seja, a gente pode ter canais que só recebem informação, canais que só
// enviam informação, canais nos quais eu posso colocar informação e canais dos quais eu só posso tirar
// informação.
// Quando eu configuro um canal como sendo send channel ou receive channel eles são tipos diferentes igual
// um int é uma coisa e um int32 é outra e um int64 é outra ainda, então um canal send e um canal receive
// são coisas diferentes.
// Isso permite que os type-checking mechanisms, ou seja, os mecanismos de checagem de tipo do compilador
// façam com que não seja possível, por exemplo, escrever num canal de leitura. Um outro exemplo seria você
// tentar tirar gasolina do tanque de um carro, seria mais fácil você colocar gasolina, então para evitar
// problemas desse tipo a gente usa canais send e canais receive.
// Para criar um send channel você vai utilizar essa anotação "chan<-", colocar informação no canal.
// Para criar um receive channel você vai utilizar essa anotação "<-chan", retirar informação do canal.
// Assignment/conversion de canais: no nosso exemplo criamos um canal bidirecional dentro da função send()
// esse canal se torna um canal send, dentro da função receive() esse canal se torna um canal receive, então
// eu posso pegar um canal geral (bidirecional) e transformar ele num canal específico, ou seja, um canal send
// ou um canal receive.
// Converter um canal geral para um canal específico é possível.
// Se eu tenho um canal específico e quero atribuir um outro canal específico não é possível.
// Se eu tenho um canal específico e quero converter para um canal geral não é possível.
// Não consigo fazer atribuição de tipos diferentes, se eu tenho um canal específico e quero atribuir esse
// valor para um canal geral não é possível porque são tipos diferentes.
