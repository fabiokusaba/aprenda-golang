package main

import "fmt"

var y = "olá bom dia"

func main() {
	x := 10

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
}

// O operador curto de declaração é representado por ":=" (gopher) e o que ele faz? Ele serve para declarar variáveis, ou seja,
// ele serve para atribuir valor a variáveis? Não, ele serve para declarar variáveis, declarar variável é você dizer para o
// programa: "olha, a gente tem uma variável nova, ela tem o nome tal e ela tem o tipo tal", isso é declarar uma variável.
// Atribuir valor a uma variável, por exemplo x = 10, é uma atribuição que quer dizer outra coisa e para isso utilizamos um
// operador diferente que é o operador de atribuição.
// A marmota tem tipagem automática, ou seja, eu posso declarar uma variável e o operador curto vai atribuir o tipo com base no
// valor da variável. O operador curto de declaração automaticamente define o tipo da variável.
// Só posso repetir o operador curto de declaração quando houverem declarações, o que é uma declaração? É dizer para o programa
// "olha, aqui tem uma variável nova", palavra-chave "nova", ou seja, se eu quiser atribuir um valor novo para x, é uma variável
// que já existe, eu vou utilizar o operador de atribuição "=". Se eu tentar utilizar um operador de declaração, marmota, vou ter
// um erro porque estou tentando declarar variáveis no contexto em que eu só posso fazer uma atribuição porque essa variável já
// existe.
// No exemplo onde eu tenho uma atribuição no x, mas uma declaração no z o operador curto de declaração funcionou, então disso a
// gente tira que ele funciona assim: ele deve ser usado para declarar pelo menos uma variável, ele pode atribuir outras, mas ele
// vai declarar pelo menos uma variável e se eu não for declarar nenhuma variável eu não posso usar o operador de declaração.
// Ele é diferente do operador de atribuição, a gente tem um operador que atribui e outro operador que declara. O operador curto de
// declaração ele também atribui como em nosso exemplo onde estamos declarando uma variável nova, x, e atribuindo um valor, 10, para
// ela ao mesmo tempo, mas ele não pode fazer só a atribuição ele tem que fazer declaração também.
// A marmota só funciona dentro de codeblocks e o que é isso? As variáveis x e z só vão funcionar dentro desse codeblock, ou seja,
// dentro da nossa função main, isso se chama scope e scope quer dizer abrangência e abrangência determina aonde que a variável vai
// estar visível ou acessível dentro de um programa. A variável y tem package level scope isso quer dizer abrangência ao nível do
// pacote, ou seja, essa variável y vai estar disponível em qualquer lugar dentro desse package enquanto que a variável x só vai estar
// disponível dentro desse codeblock mais especificamente da sua declaração para frente.
// A marmota, operador curto de declaração, só funciona dentro de um codeblock isso quer dizer que eu não posso usar a marmota aqui
// fora porque isso vai dar erro, ela necessariamente tem que ser utilizada dentro de um codeblock se você quer fazer uma declaração com
// package level scope você precisa usar o "var" e colocá-lo fora de um codeblock.
// Terminologia -> keywords (palavras-chave) são termos reservados da linguagem que não podem ser utilizados como identificadores,
// identificadores é o nome que você dá para identificar uma constante, uma variável, uma função, um método, uma interface, etc.
// As palavras-chaves (keywords) somente podem ser usadas pra função específica delas, por exemplo você pode utilizar a palavra package
// para determinar o nome do seu package, mas você não pode utilizar pra mais nada.
// Operadores e operandos -> toda expressão é feita de operadores e operandos, exemplo 10 + 10, o sinal de adição é o operador e os números
// que são os operandos.
// Expressão -> é qualquer coisa que produz um resultado, por exemplo 10 + 10, é um resultado, produz um resultado, tem um valor. Ou eu
// posso dizer 10 == 10, isso é uma expressão e ela vai me retornar um bool, um tipo que pode ser verdadeiro ou falso, sim ou não, 1 e 0,
// em homenagem a George Boole. Ou seja, uma expressão é qualquer coisa que produza um resultado, mas que não gera uma ação.
// statement (declaração, afirmação) -> é uma linha de código, geralmente é uma linha mas não necessariamente vai ser, é uma instrução que
// forma uma ação e ela é formada de expressões. Ou seja, nesse exemplo x := 10 < 20, a ação é pegar o resultado dessa expressão e salvar
// como uma variável e já que ele é o operador curto ele faz tanto a declaração quanto a atribuição.
// Distinção entre expressão e statement -> um statement é formado de expressões de uma ou mais expressões, uma expressão ela não é um
// statement. O que a gente tem que saber disso é que quando estamos falando de variáveis a gente vai usar declaração no sentido de
// declaration, ou seja, no sentido de dizer para o sistema que isso é uma variável nova e quando a gente tiver falando de instruções linha
// por linha a gente vai estar falando de declaração no sentido de statement, ou seja, ações que o nosso programa vai fazer.
