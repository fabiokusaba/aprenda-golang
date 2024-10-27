package main

import "fmt"

func main() {
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	fmt.Println(umaslice)

	umaslice = append(umaslice, 5, 6, 7, 8)

	fmt.Println(umaslice)

	// Os três pontinhos estão dizendo que ao invés de pegar a fatia de int que é "outraslice", é para pegar o conteúdo
	// que está dentro dessa fatia, os itens dessa fatia. Ele é chamado de unfurl e na documentação oficial está como
	// enumeration.
	umaslice = append(umaslice, outraslice...)
	fmt.Println(umaslice)
}

// A função append serve para anexar itens a uma slice e ela faz parte do package builtin.
// Ela toma dois elementos, no caso ele toma uma slice e uma quantidade variádica de elementos do tipo da slice.
