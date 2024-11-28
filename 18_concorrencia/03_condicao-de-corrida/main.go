package main

func main() {

}

// Não se comunique através de compartilhar memória, ao invés disso, compartilhe memória para se comunicar.
// Programar de maneira concorrente em muitos ambientes é difícil por causa das sutilezas necessárias para
// implementar o acesso as variáveis compartilhadas, ou seja, você tem várias goroutines e elas vão acessar
// uma mesma variável isso vai dar zebra.

// Em Go a gente tem uma abordagem diferente os valores que a gente quer compartilhar são passados de um
// lugar para o outro através de canais e esses valores nunca são compartilhados por duas coisas ao mesmo
// tempo.

// Somente uma goroutine tem acesso ao valor em qualquer dado momento e condições de corrida não podem
// acontecer por design, por definição porque a linguagem utilizando canais é impossível você fazer uma
// condição de corrida.

// Uma condição de corrida é uma falha no sistema ou processo em que o resultado do processo é inesperadamente
// dependente da sequência ou sincronia de outros eventos.
// O que é yield? É quando, por exemplo você tem lá o seu computador de 2005 que só tem um processador e
// você está executando o browser e o programa que toca música e a cada fração de microssegundo lá o processador
// alterna entre um processo e outro, esse momento em que executei um pouquinho desse programa quando chegou o
// momento em que vou trocar de execução esse é o yield, é quando eu digo "tá, agora deu disso aqui e vamos pra
// outra coisa." Essa troca de processos se chama yield.
