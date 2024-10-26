package main

import "fmt"

func main() {
	idade := 16

	if idade >= 18 {
		fmt.Println("você é maior de idade")
	} else if idade >= 15 {
		fmt.Println("você é adolescente")
	} else {
		fmt.Println("você é menor de idade")
	}
}
