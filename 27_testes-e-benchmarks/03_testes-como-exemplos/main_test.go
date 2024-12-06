package main

import (
	funcionalidades "aprenda-golang/27_testes-e-benchmarks/02_testes-em-tabela"
	"fmt"
)

func ExampleSoma() {
	fmt.Println(funcionalidades.Soma(4, 2, 1))
	fmt.Println(funcionalidades.Soma(4, 2, 2))
	fmt.Println(funcionalidades.Soma(4, 2, 3))
	// Output:
	// 7
	// 8
	// 9
}

// Outra maneira é fazer testes como exemplos.
// Eles são testes, mas eles também são os mesmos exemplos que aparecem na documentação.
// Eu escrevo uma documentação e nessa documentação eu digo que quando eu rodar esse código o resultado vai ser esse.
// Então, isso tem dois efeitos: um ele roda esse teste confere se realmente quando rodar aquele código o resultado vai
// ser aquele. E dois eu estou gerando documentação, então os meus usuários vão poder ler essa documentação e utilizar
// essa documentação pra poder utilizar o meu código.
// Para exemplos o formado é "func ExampleFuncao()"
// Deve haver um comentário "// Output: resultado", que é o que será testado.
