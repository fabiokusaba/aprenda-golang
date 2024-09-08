package main

import "fmt"

func main() {
	numerodebytes, erros := fmt.Println("Hello world!", "Oi galera!", 100)
	fmt.Println(numerodebytes, erros)
}

// O programa "Hello world!" é utilizado em todas as linguagens porque ele demonstra os elementos básicos de uma linguagem.
// package main -> quando estamos fazendo um programa você não vai colocar tudo no mesmo arquivo de texto, banco de dados,
// imagens, servidores, etc. Você vai separar para manter uma organização, mas aí você tem um monte de código, um monte de
// arquivos, um monte de packages e como é que o computador sabe aonde tudo começa? Ele não sabe, então a gente tem que dizer
// pra ele e a gente diz pra ele denominando algum package no nosso programa de "package main".
// main é um termo em inglês que significa principal, então a gente está dizendo que esse é o package principal.
// A mesma coisa no "function main" porque num programa a gente pode ter várias funções, várias declarações, métodos, interfaces
// constantes, etc. Um monte de coisa e como é que o programa sabe aonde ele deve começar? Novamente ele não sabe, então a gente
// diz pra ele: "function main" aqui é onde você começa, aqui é onde você termina. Isso quer dizer que a primeira linha da função
// "main" vai ser a primeira linha executada do seu programa e ao finalizar a execução da última instrução da sua função "main" o
// seu programa finaliza e encerra, ou seja, a função "main" é a porta de entrada e a porta de saída de qualquer programa que você
// escrever em go.
// Uma coisa pra gente olhar também é a notação dos pacotes, é importante a gente prestar atenção nessa notação: "pacote.Identificador"
// tem muitas coisas na linguagem que mudam, por exemplo pra você declarar uma variável você pode fazer isso de várias maneiras, mas
// essa notação dos pacotes não muda nunca, ela vai ser sempre a mesma por isso devemos prestar atenção.
// fmt.Println -> exibe o valor na tela, entre cada frase vai ter um espaço. A função além de fazer a ação associada a ela, que é
// colocar as informações na tela, ela vai nos dar um retorno um dizendo com quantos bytes ela trabalhou e a outra dizendo quais erros
// houveram caso tenham havido erros. Como faço para acessar isso? Eu preciso botar o retorno de uma função dentro de variáveis.
// Em Go não podemos colocar a informação dentro das variáveis e não utilizá-las depois.
// ...interface{} -> esses três pontinhos indicam uma função variádica. Uma função variádica é uma função que trabalha com qualquer
// número de argumentos e consegue lidar com eles. Todos os tipos em Go implementam a interface vazia, então isso quer dizer que eu
// posso colocar quantos elementos eu quiser de qualquer tipo que eu quiser.
// Ou seja, na declaração de uma função eu tenho o nome, o que ela recebe e o que ela me entrega, ou seja, o que a função me dá que é
// o retorno.
// Variáveis não utilizadas -> em Go a maneira idiomática de se usar a linguagem é ser sempre eficiente, sempre fazer as coisas com
// propósito, não deixar sujeira, não deixar distrações no meio do nosso programa e uma variável que está declarada e não utilizada
// a linguagem interpreta como sujeira, como ineficiência, então ela não permite que a gente faça isso porque vai contra os seus
// princípios.
// Quando a gente tem uma variável que a gente não vai utilizar a gente descarta esse valor e fazemos isso colocando um underline "_"
// na linguagem ele é chamado de "Blank identifier". Isso quer dizer que estou recebendo informações e não quero utilizar essas
// informações.
// Uma variável é um objeto, uma posição com endereço na memória capaz de reter ou representar um valor ou uma expressão.
