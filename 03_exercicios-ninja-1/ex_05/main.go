package main

import "fmt"

type meutipo int

var x meutipo

var y int

func main() {
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
