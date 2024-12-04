package main

func main() {

}

// Para quem programa em outras linguagens já deve ter visto try-catch-finally em Go não usamos isso.
// Por que Go não tem exceções? A gente acredita que juntar exceções no nosso controle do programa e ficar usando
// idiomas como try-catch-finally resulta em código bagunçado, isso também encoraja programadores a chamar muitas coisas
// comuns como por exemplo não conseguir abrir um arquivo de exceções quando não são exceções são uma coisa que acontece
// toda hora.
// Então, o Go tem uma abordagem diferente a gente vai usar sempre aquele múltiplos retornos das funções, ou seja, se a
// sua função retorna um valor como o erro vai funcionar? A sua função vai te retornar um valor e um erro, se esse erro
// for vazio a gente continua porque deu tudo certo, se esse erro tiver algum valor temos que interromper o nosso
// programa e de alguma maneira dizer para aí que deu coisa errada, essa é a maneira que é feita em Go e que os creators
// da linguagem acreditam ser a mais simples.
// A linguagem Go tem algumas funções pré prontas pra sinalizar e se recuperar de erros realmente excepcionais. Existem
// mecanismos para lidar com situações realmente excepcionais, que na opinião deles são realmente exceções, existem
// mecanismos para se recuperar de quando ocorrem essas exceções que é o que aconteceria com o try-catch-finally que
// muitas linguagens usam, mas que na opinião deles é diferente, é separado, uma coisa é um erro simples como a gente
// tem que abrir um arquivo e ele não foi encontrado, outra coisa é realmente uma exceção, uma coisa catastrófica, então
// a gente tem que lidar com essas coisas de maneira diferente.
// O tipo error é uma interface pré pronta da linguagem, é a interface convencional para representar erros e um valor
// vazio representa que não há erros.
// A gente só precisa de um metodo para implementar essa interface Error() e ele tem que retornar uma string, ou seja,
// ele tem que retornar a mensagem de erro.
// O package errors implementa funções para manipular erros.
// Um detalhe importante é que é interessante criar o hábito de lidar com erros imediatamente, similar a defer close,
// ou seja, lembra quando estudamos a função defer que quando a gente abre um arquivo a gente imediatamente fecha o
// arquivo só que a gente coloca um defer, então a gente vai abrir o arquivo fazer toda a nossa função e antes da nossa
// função finalizar ele vai rodar o defer que é fechar o arquivo, mas quando sempre eu crio o hábito de abrir o arquivo
// e fechar logo em seguida, mesmo com o defer, o nosso código fica mais limpo porque tudo o que estou abrindo estou
// fechando.
// Então, aqui é a mesma coisa para lidar com erros, lide com erros imediatamente, então se você tem uma função que te
// retorna um valor e um erro imediatamente na linha de baixo já verifica se tem algum erro e já lida com ele, não deixe
// pra verificar o erro depois.
