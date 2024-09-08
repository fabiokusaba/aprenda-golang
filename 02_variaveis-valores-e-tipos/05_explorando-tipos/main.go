package main

import "fmt"

var x int

func main() {
	x = 10
	fmt.Printf("%v, %T", x, x)
}

// Tipos em Go são extremamente importantes.
// Tipos em Go são estáticos, ou seja, se eu declarei uma variável x como sendo um integer ela vai ser um integer até o fim
// dos tempos, até acabar o programa, o tipo é pra sempre se você precisar de outro tipo você vai precisar de outra variável
// é impossível mudar o tipo de uma variável uma vez que esse tipo foi declarado.
// Ou seja, Go is a static type language, é uma linguagem de tipagem estática.
// Ao declarar uma variável essa variável só poderá conter valor do tipo que eu declarei quando declarei a variável.
// O tipo pode ser deduzido pelo compilador. Como já vimos quando utilizamos a marmota ela poẽ o tipo sozinha e quando a gente
// usar o var se a gente não coloca o tipo explicitamente ela também deduz o tipo sozinha.
// O valor pode ser declarado explicitamente.
// Um fator importante é que a declaração com package level scope caso a gente não atribua um valor no momento da declaração esse
// valor só vai poder ser atribuído dentro de um codeblock.
// Os tipos básicos na linguagem Go são: int, string e bool.
// Além dos tipos de dados primitivos a gente tem os tipos compostos e o que são eles? São tipos feitos pelo usuário formados por
// tipos primitivos. Por exemplo eu posso ter um tipo que vai ter dez strings, um slice, ou eu posso ter um map com valores de
// vários tipos dentro, ou eu posso ter um struct com maps, slices, arrays, com qualquer coisa que eu quiser dentro dele, esses são
// os tipos de dados compostos: slice, array, struct, map.
// O ato de definir, criar, estruturar tipos compostos chama-se composição.
