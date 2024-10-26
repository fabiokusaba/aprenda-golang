package main

import "fmt"

func main() {
	anoemquenasci := 1995
	anoateoqualquerocontar := 2088

	for {
		if anoemquenasci > anoateoqualquerocontar {
			break
		}

		fmt.Println(anoemquenasci)
		anoemquenasci++
	}
}
