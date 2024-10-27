package main

import "fmt"

func main() {
	a := closure()
	b := closure()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println("")

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func closure() func() int {
	x := 0
	return func() int {
		x += 10
		return x
	}
}
