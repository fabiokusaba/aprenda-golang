package main

import "fmt"

func main() {
	//                  0.          1.           2.         3.                4.
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatro queijos", "marguerita"}

	//fatia := sabores[0:2]
	//fatia := sabores[3:5]
	//fatia := sabores[2:5]
	//fatia := sabores[2:len(sabores)]
	//fatia := sabores[2:]
	//fatia := sabores[:2]
	fatia := sabores[:]
	fmt.Println(fatia)

	fmt.Println(sabores[0], sabores[1], sabores[2], "\n")

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

	// Como deletar um item de uma slice, basicamente seleciono tudo o que tem antes, seleciono tudo o que tem depois
	// do item que eu quero excluir, ou seja, se eu quero excluir a terceira hora do relógio eu faço uma seleção do 0
	// até o 3, depois do 4 até o 12, ou seja, a hora 3 ficou fora da minha seleção, então simplesmente junto aquelas
	// duas fatias e aquela hora vai ser excluída.
	// Então, para deletar o valor de uma slice eu simplesmente crio uma slice nova, no nosso caso estamos reatribuindo
	// para esse identificador "sabores" um outro valor, uma outra slice, essa slice vai se formar de uma slice que é
	// ["pepperoni, "mozzarela"] e de outra slice que é ["quatro queijos", "marguerita"], ou seja, o "abacaxi" está fora
	sabores = append(sabores[:2], sabores[3:]...)
	fmt.Println(sabores)
}

// Para fatiarmos uma fatia podemos fazer da seguinte forma: x[:], x[a:], x[:b], x[a:b].
// O "a" é primeiro item e ele é incluso.
// E o "b" não é incluso.
