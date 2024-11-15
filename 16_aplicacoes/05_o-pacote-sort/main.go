package main

import (
	"fmt"
	"sort"
)

func main() {
	ss := []string{"abóbora", "maçã", "laranja", "beringela", "berinjela"}
	fmt.Println(ss)

	// Ordenando alfabeticamente
	sort.Strings(ss)
	fmt.Println(ss)

	si := []int{987, 123, 789, 654, 321, 456}
	fmt.Println(si)

	// Ordenando de forma crescente
	sort.Ints(si)
	fmt.Println(si)
}

// O pacote sort é um pacote de ordenação pra ordenar elementos de slices.
// O pacote sort ele nos dá elementos primitivos para ordenar slices e coleções definidas pelo usuário.
