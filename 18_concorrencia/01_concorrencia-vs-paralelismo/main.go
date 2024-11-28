package main

func main() {

}

// Concorrência: é um padrão de design que faz com que vários elementos sejam executados de maneira independente e
// com que um não dependa do outro, ou seja, a gente tem várias coisas acontecendo ao mesmo tempo e uma coisa não
// tem que necessariamente esperar a outra para que o processamento possa continuar.

// Qual a diferença entre concorrência e paralelismo? A concorrência é uma maneira de pensar, é um jeito de você
// organizar o seu programa. O paralelismo vem quando a gente tem código concorrente sendo executado num sistema
// que possui vários processadores.

// Na programação a gente não quer saber de paralelismo, a gente quer saber de concorrência, a gente vai fazer um
// código com padrões concorrentes, a gente vai criar funções concorrentes, se isso vai ser executado de maneira
// paralela ou não vai depender da arquitetura aonde você está executando o seu código, vai depender do processador
// do seu computador, por padrão a linguagem Go executa código Go utilizando o máximo de núcleos possíveis, então
// o paralelismo vai ocorrer sempre que possível e esse problema é do runtime da linguagem sendo que o nosso
// problema é escrever código concorrente.
