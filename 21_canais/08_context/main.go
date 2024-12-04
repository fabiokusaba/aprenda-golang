package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("--------------------------")

	ctx, cancel := context.WithCancel(ctx)
	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("--------------------------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("--------------------------")
}

// Para que serve context? Imagine o seguinte: que você tenha um site tipo Netflix aí o seu cliente digitou o nome de um
// programa de tv lá no search só que o seu banco de dados é meio fragmentado então você tem 50 banco de dados diferente
// um pra séries, um pra filmes, etc... Então o usuário digitou um nome lá e deu enter e você despachou lá uma porrada
// de goroutines cada uma pra puxar informação de uma fonte diferente e a medida que essas goroutines vão completando a
// sua execução os resultados vão aparecendo na tela da pessoa.
// Isso é legal porque ao invés de você ir um por um buscando informações você faz várias ao mesmo tempo, mas imagina
// que o cara já sabia do negócio que ele queria, ele deu enter e o primeiro resultado que apareceu na tela ele já click
// então a pessoa não está mais naquela página de busca, não quer mais saber daqueles resultados, mas suponhamos que
// você ainda tem lá 40 goroutines buscando informação de outros bancos de dados que não responderam ainda, o que seria
// legal você fazer? Seria legal você poder cancelar essas goroutines, matar elas porque essa informação já não é mais
// relevante, essa tarefa que você mandou elas fazerem não tem mais necessidade então você pode matar elas ao invés de
// ficar torrando memória e processamento com isso você usa esses recursos para outra coisa ou economiza uma grana.
// E esse é um dos motivos que a gente tem o package Context que é pra gente poder trocar mensagens entre goroutines
// com esse objetivo.
// Destaques da documentação que você precisar prestar atenção!!!
// context.Background -> basicamente ele é a base do Context, então se você quer usar um context pra qualquer coisa você
// precisa de um Background.
// context.WithCancel -> Cancel é o que comentamos, sua goroutine está lá fazendo alguma coisa e você diz pra ela chega,
// fora, se joga pela janela porque não preciso mais de ti, ou seja, você vai criar um contexto com Background, você vai
// usar esse contexto para fazer um contexto com Cancel e isso vai te dar não só um contexto que é o que você tinha com
// Background mas você vai ter também a função cancel e a função cancel manda um sinal de cancelamento pra aquela
// goroutine.
// Na goroutine você pode ter um select e um dos cases pode ser ctx.Done(), o que acontece aqui? Acontece que quando
// você executa o cancel isso aqui é enviado para a sua goroutine e aí você faz um return e você mata aquela goroutine
// ou seja, se você tiver um select com esse case e ele for o case que está rodando acabou aquela goroutine, se não
// continua o que estava fazendo.
// Tem algumas coisas que a gente pode verificar que são os erros do contexto ctx.Err().
// Além do WithCancel tem o WithDeadline e o WithTimeout que basicamente: o Timeout é um timer onde você diz que essa
// goroutine tem um minuto para executar e passou esse tempo ela morre, e o Deadline também é um temporizador mas você
// dá a hora exata então essa goroutine tem até as 17h para executar, deu esse horário ela morreu.
