package main

import "fmt"

func main() {
	x := 10
	if x > 100 {
		fmt.Println("chis é maior que cem")
	} else if x < 10 {
		fmt.Println("chis é menor que déis")
	} else {
		fmt.Println("chis não é menor que déis nem maior que cem")
	}
}

// Agora vamos falar das extensões do if que são else if e else.
// O que significa if e else? Significa assim caso alguma coisa faça isso, caso contrário (else) faça outra coisa.
// Ou eu posso dizer caso uma coisa faça isso, caso outra coisa (else if) faça isso, caso contrário, ou seja, caso nenhuma das
// anteriores (else) faça isso.
// Ou eu posso dizer caso uma coisa faça isso, caso outra faça outra, caso outra faça outra, etc etc etc, até o final caso nenhuma
// delas faça a última.
