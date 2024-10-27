package main

import "fmt"

func main() {
	dogs := [][]string{
		[]string{
			"Luna",
			"Luninha",
			"Fazer arte",
		},
		[]string{
			"Luninha",
			"Luna",
			"Aprontar todas",
		},
		[]string{
			"Luna",
			"Luna",
			"Latir sem parar",
		},
	}

	for _, v := range dogs {
		for _, item := range v {
			fmt.Println(item)
		}
	}

}
