package main

import "fmt"

func main() {
	//si := []int{10, 10, 1, 2, 3, 5}
	//total := soma(si...)

	total := soma()

	fmt.Println(total)

}

// Quando temos uma slice podemos passar os elementos individuais através deste "..." operador.
// Numa função variádica eu posso passar uma quantidade ilimitada ou zero valores que não teremos problema.
func soma(x ...int) int {
	soma := 0

	for _, v := range x {
		soma += v
	}

	return soma
}
