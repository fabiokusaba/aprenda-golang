package main

import "fmt"

func main() {
	amigos := map[string]int{
		// Keys: values
		"alfredo": 5551234,
		"joana":   9996674,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	// Adicionando valores no map.
	amigos["gopher"] = 444444

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"])

	// "comma ok" idiom -> as vezes você precisa distinguir uma entrada que não existe de um valor zero e como posso
	// fazer isso? A expressão "vírgula ok", então ao invés de eu simplesmente pegar uma variável e adicionar um valor
	// pra ela do map eu uso duas variáveis: a variável com o nome que eu quero e "ok", e nesse "ok" eu vou ter um bool
	// que vai dizer true caso exista um valor e false caso não exista.
	if sera, ok := amigos["fantasma"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(sera)
	}

}

// Maps são uma estrutura de dados do tipo chave e valor ou key value pair (pares chave e valor), a maneira mais fácil
// de ilustrar seria a sua agenda de telefone em que você tem o nome e o telefone, ou seja, nome e valor.
// A diferença entre a sua agenda e os maps é que os maps não tem ordem, não são ordenados, portanto se você fizer um
// range em um map a ordem vai ser aleatório, mas em contraste o lookup dos maps tem a performance extremamente boa
// então você coloca qualquer nome ali que ele vai conseguir achar bem rápido.
