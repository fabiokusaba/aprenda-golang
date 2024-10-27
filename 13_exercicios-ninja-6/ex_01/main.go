package main

import "fmt"

func main() {
	r := retornaInt(7)
	n, s := retornaIntString(7, "Fabio")
	fmt.Println(r)
	fmt.Println(n, s)
}

func retornaInt(x int) int {
	return x
}

func retornaIntString(x int, y string) (int, string) {
	return x, y
}
