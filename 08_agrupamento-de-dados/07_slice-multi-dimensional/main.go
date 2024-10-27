package main

import "fmt"

func main() {
	ss := [][]int{
		//Índice: 0  1  2    Índice:
		[]int{1, 2, 3}, // 0
		[]int{4, 5, 6}, // 1
		[]int{7, 8, 9}, // 2
	}

	fmt.Println(ss)
	fmt.Println(ss[1])
	fmt.Println(ss[1][1])
}

// Slice multi-dimensional você tem que enxergar como sendo uma planilha no excel.
// Vimos que o slice é uma coleção de valores em sequência, por exemplo 1, 2, 3, 4.
// Um slice multi-dimensional vai conter vários slices abaixo do outro como se fosse uma planilha.
// Isso funciona como sendo um slice of slices, ou seja, uma fatia de fatias.
