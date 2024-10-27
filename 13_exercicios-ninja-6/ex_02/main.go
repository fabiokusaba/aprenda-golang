package main

import "fmt"

func main() {
	si := []int{1, 2, 3, 4, 5}
	fmt.Println(somaVariadica(si...))
	fmt.Println(somaSliceOfInt(si))
}

func somaVariadica(x ...int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}

func somaSliceOfInt(x []int) int {
	s := 0
	for _, v := range x {
		s += v
	}
	return s
}
