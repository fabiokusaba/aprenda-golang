package main

import (
	"fmt"
	"sync"
	"time"
)

// Declarando o nosso WaitGroup em package level scope para que ele se torne global e possa ser utilizado
// na nossa aplicação
var wg sync.WaitGroup

func main() {
	// Aqui temos duas funções sendo executadas de maneira linear, ou seja, uma primeiro e depois a outra.
	//func1()
	//func2()

	// Nós queremos que essas funções sejam executadas de maneira concorrente, ou seja, queremos que as duas
	// sejam executadas ao mesmo tempo, e como fazemos isso? Utilizando Goroutines.
	// O que são Goroutines? São "threads".
	// Threads: linha ou encadeamento de execução, é uma forma de um processo dividir a si mesmo em duas ou
	// mais tarefas que podem ser executadas concorrencialmente.
	// Goroutines são linhas ou encadeamento de função, é como se cada goroutine que eu fizer fosse executada
	// independentemente, é outro processo, outro programa, por exemplo a gente pode ter o browser e o programa
	// de música aberto ao mesmo tempo, então dois goroutines é como se fosse isso: dois programas separados
	// cada um fazendo as suas coisas, a princípio independentes.
	// Quando colocamos a palavra go na frente de uma função ela se torna uma goroutine, ou seja, ela vai ser
	// executada independentemente e como a gente tem que pensar nisso? É assim, imagina que você pegou essa
	// função colocou na agulha e deu um tiro, a partir do momento em que você apertou o gatilho você não tem
	// mais como mexer naquela bala.
	// sync.WaitGroup: é um grupo de espera, ou seja, na minha função main eu vou dizer para o programa esse
	// programa vai ter 20 goroutines, ou seja, 20 foguetes espaciais que vão estar fazendo suas coisas a gente
	// não vai saber o que eles estão fazendo, mas temos que esperar por eles então vou ter 20 threads que eu
	// quero esperar.
	// No WaitGroup a gente vai ter três funções principais: a função Add vou estar dizendo para o WaitGroup
	// adiciona um ou vou dizer adiciona 20, vou dizer quantas coisas eu quero que ele espere, por quantas coisas
	// eu quero que ele espere. A função Done vou utilizar dentro de cada goroutine, ou seja, vou estar mandando
	// um sinal para o meu WaitGroup da função main dizendo que a função já acabou. E a função Wait é o que eu
	// vou utilizar pra dizer para o meu programa "pera aí cara, não encerra ainda porque a gente tem que esperar
	// as goroutines."
	wg.Add(2) // dizendo ao programa que tenho 1 goroutine para ele esperar.
	// add(total de funções)
	go func1()
	func2()
	// espera!
	wg.Wait() // esperar o processamento da goroutine
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
		time.Sleep(1 * time.Second)
	}
	// deu! terminou de executar
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}
