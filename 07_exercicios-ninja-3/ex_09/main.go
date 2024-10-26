package main

import "fmt"

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Brasil, o país do futebol")
	case "starcraft":
		fmt.Println("Coreia, viciados por e-sports")
	case "basquete":
		fmt.Println("NBA, astros do basquete")
	}
}
