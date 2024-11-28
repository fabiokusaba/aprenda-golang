package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Contando a quantidade de processadores que a gente tem
	fmt.Println("CPUs:", runtime.NumCPU())

	// Aqui eu vou ver quantos goroutines eu tenho
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	// Declarando uma variável em comum que todos irão acessar.
	contador := 0

	// Declarando o total de goroutines que vamos ter
	totaldegoroutines := 1000

	// Vou precisar de um WaitGroup para evitar que o meu programa feche antes que as minhas goroutines
	// executarem na prática, ou seja, só quero sair do programa quando elas terminarem de executar.
	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	// Vamos fazer uma função para disparar essas goroutines
	for i := 0; i < totaldegoroutines; i++ {
		// Aqui vai ser onde eu vou lançar as minhas funções
		go func() {
			// Vou ter uma variável interna que vai ser o valor do nosso contador externo
			v := contador
			// yield
			runtime.Gosched()
			// Incrementando o valor da variável
			v++
			// Salvando no meu contador esse valor incrementado da minha variável
			contador = v
			// Aqui ao final vou precisar de um Done para dizer que essa goroutine acabou
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	// Aqui eu preciso dizer: "pera aí, não termina a execução ainda", então quando chegar aqui a minha
	// função main vai ficar esperando até as minhas goroutines dizerem que terminaram
	wg.Wait()

	fmt.Println("Goroutines:", runtime.NumGoroutine())

	// E aqui vou mostrar o valor do contador
	fmt.Println("Valor final:", contador)
}

// Isso aqui quer dizer que várias go funcs, várias threads acessaram a minha variável contador ao mesmo
// tempo e deu aquele problema que tínhamos discutido, ou seja, várias leram ao mesmo tempo, incrementaram
// a partir do mesmo valor básico e quando elas salvaram não foram dez somas sequenciais
