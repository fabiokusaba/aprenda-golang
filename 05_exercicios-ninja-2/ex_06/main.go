package main

import "fmt"

const (
	_ = 2024 + iota
	ano1
	ano2
	ano3
	ano4
)

func main() {
	fmt.Println(ano1)
	fmt.Println(ano2)
	fmt.Println(ano3)
	fmt.Println(ano4)
}
