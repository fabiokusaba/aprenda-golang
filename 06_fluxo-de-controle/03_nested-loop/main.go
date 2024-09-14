package main

import "fmt"

func main() {
	for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês ", mes)
		for dias := 1; dias < 60; dias++ {
			fmt.Println("Dia ", dias, ", ")
		}
		fmt.Println()
	}
}

// Nested loop ou repetição hierárquica -> por exemplo temos três ponteiros no relógio e eles sempre repetem, ou seja, estão
// sempre girando só que a gente tem repetições que ocorrem dentro de repetições, por exemplo cada vez que o ponteiro dos minutos
// gira o ponteiro dos segundos deu várias voltas completas, cada vez que o ponteiro das horas dá uma volta completa o ponteiro dos
// minutos deu um monte de voltas completas, ou seja, eu tenho uma repetição dentro de outra repetição dentro de outra repetição
// mais lenta.
// E a mesma coisa no calendário você tem as horas dentro do dia sempre repetindo um dia depois do outro, você tem os dias dentro da
// semana, você tem as semanas dentro dos meses, você tem os meses dentro dos anos e por aí vai. Ou seja, repetições menores dentro
// de repetições maiores, ciclos pequenos dentro de ciclos grandes.
