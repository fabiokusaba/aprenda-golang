package main

import "fmt"

var x interface{}

func main() {
	x = true

	switch x.(type) {
	case int:
		fmt.Println("é um int")
	case bool:
		fmt.Println("é um bool")
	case string:
		fmt.Println("é um string")
	case float64:
		fmt.Println("é um float")
	}

	switch y := 2; {
	case y == 1:
		fmt.Println("é um 1")
	case y == 2:
		fmt.Println("é um 2")
	case y == 3:
		fmt.Println("é um 3")
	case y == 4:
		fmt.Println("é um 4")
	}
}
