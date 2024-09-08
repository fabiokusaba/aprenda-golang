package main

import "fmt"

// Esse aqui não, ele vai ficar indefinido até ele ser usado.
const x = 10

// Declaração de múltiplas constantes
const (
	z = 30
	w = 40
	v = 50
)

// No momento que eu executar esse programa, no momento em que o compilador ler essa linha ele já vai decidir isso é um int.
var y = 10

var c int
var d float64

func main() {
	// A constante x se torna um int a partir do momento que eu mandei ele interagir com a variável c que é um int.
	// c = x
	c = y
	fmt.Println(c)

	// A constante x não é nem int e nem float, ele não é nada e só vai se tornar alguma coisa quando ele for usado.
	d = x
	fmt.Println(d)

	fmt.Println(z, w, v)
}

// Constantes são como se fossem variáveis só que são valores imutáveis, você não pode mudar nunca. As vezes a gente tem valores
// no nosso programa que vão ser sempre a mesma coisa, que não vão mudar então eles a gente põe em constantes.
// O endereço do seu banco de dados em seu app pode ser uma constante porque ele não vai mudar ao longo da execução do app, ele vai
// ser um valor que você vai colocar no momento da compilação, você compila o programa e aquele valor não vai mudar durante a execução.
// As constantes podem ser tipadas ou não.
// Uma coisa importante das constantes é que as constantes não tipadas elas só terão tipo atribuído a elas quando elas forem usadas. Por
// exemplo qual o tipo de 42? int? uint? float64? Ou seja, é uma flexibilidade conveniente.
// A gente chega a conclusão que essa nossa constante x pode ser usada tanto com int quanto com float e o y não, o y só pode ser usado
// como int. Reiterando, o tipo de uma constante só é definido no momento do uso, o tipo de uma variável é definido no momento da
// atribuição.
