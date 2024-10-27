package main

import "fmt"

func main() {
	slice := make([]int, 5, 10)

	slice[0], slice[1], slice[2], slice[3] = 1, 2, 3, 4

	slice[4] = 5

	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 10)

	fmt.Println(slice, len(slice), cap(slice))

	slice = append(slice, 11)

	fmt.Println(slice, len(slice), cap(slice))

}

// Slices são feitos de arrays.
// As slices são dinâmicas, diferente dos arrays elas podem mudar de tamanho.
// Sempre que uma slice muda de tamanho um novo array é criado e os dados são copiados pra lá.
// Para otimizar as coisas podemos usar o make com esse formato -> make([]T, len, cap)
// Aqui podemos ver que o make nos dá um length (len) que é o comprimento, um cap (capacity), ou seja, a capacidade
// daquele array, então posso ter um slice com 10 elementos e a capacidade de 50, o capacity se refere mais ao array
// que está por trás desse slice, por exemplo quando eu tiver 10 de length eu posso usar o append pra ir adicionando
// elementos, o array subjacente vai continuar igual com a mesma capacidade de 50, no momento em que eu tiver 50 itens
// no meu slice e eu der um append denovo aquele array vai se dissolver e vai ser criado outro com os dados, os dados
// vão ser copiados de um array para o outro e isso vai gerar um custo computacional, e aí o capacity do array novo vai
// ser maior e não mais 50.
// O make faz com que assim se você já sabe com quantos elementos você vai trabalhar antes de criar o slice você já cria
// do zero um valor "x" de gavetinhas onde você pode colocar seus dados para que o array não tenha que ser recortado,
// copiado, mudado de lugar cada vez que você adicionar um elemento.
