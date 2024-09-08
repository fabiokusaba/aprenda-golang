package main

import "fmt"

func main() {
	// Interpreted string literal
	// x := "oi bom dia\ncomo vai?\tespero \"que\" tudo bem."

	// Raw string literal
	x := `"oi bom dia\ncomo vai?\tespero \"que\" tudo bem."`
	fmt.Println(x)
}

// Strings -> existem dois tipos de strings que a gente pode declarar na linguagem Go: interpreted string literals, raw
// string literals. Um literal em ciência da computação é uma notação para representar um valor fixo no código fonte, ou
// seja, um int é um literal, um string é um literal, um bool é um literal, então para string a gente tem os interpreted
// literals e os raw literals.
// Cada caracter de uma string é chamado de rune literal.
// No package fmt temos três categorias principais:
// Print -> standart out
// Sprint -> ao invés de eu pegar esse texto que eu dei pra ele e colocar na tela ele vai pegar esse texto que eu dei pra
// ele colocar numa string e me dar essa string como retorno da minha função.
// Fprint -> é um file print, mas esse file não é necessariamente um arquivo porque em Go não tem diferença alguma eu
// escrever/colocar bytes em um arquivo ou por exemplo colocar bytes em uma conexão ao servidor ou colocar bytes em qualquer
// outro tipo de interface que aceite bytes é tudo um writer interface, ou seja, qualquer coisa que tomar como input/entrada
// um writer interface eu posso usar o Fprint para ele.
