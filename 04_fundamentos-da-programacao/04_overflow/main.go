package main

import "fmt"

func main() {
	var i uint16
	i = 65535
	fmt.Println(i)
	i++
	fmt.Println(i)
	i++
	fmt.Println(i)
}

// Overflow -> por exemplo se eu pegar esse uint16 que vai de 0 até 65535 o que acontece se eu tentar usar um valor 65536?
// Ou seja, usar um número a mais do número limite dele e o que acontece se eu tiver o número 65535 e tentar adicionar mais 1?
// Se eu tentar colocar um número maior vou ter um onverflow, isso quer dizer que ela transborda, ela passou do limite.
// Quando estou no limite e adiciono mais um o que acontece é similar ao odômetro do carro quando ele chega no seu limite 99999
// ele vira tudo 00000. Então, dependendo do que você está fazendo isso é um problema enorme.
