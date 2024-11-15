package main

func main() {

}

// A interface Writer pra gente implementá-la precisamos de um metodo Write, esse metodo vai receber uma fatia de
// bytes e ele vai retornar um número e um erro.
// No package os a gente tem um metodo que recebe um File, tipo que está definido, então a gente tem o metodo Write
// ele recebe uma fatia de bytes e ele retorna um número e um erro, ou seja, o tipo File implementa a interface
// Writer.
// No package json na função NewEncoder a gente vai ter um Writer, a gente vai receber um Writer e retornar um
// Encoder.
// Em Go para tudo a gente vai usar interfaces, as interfaces são a maneira da gente conectar partes diferentes do
// nosso código e por exemplo a interface permite que eu pegue uma tomada na parede que conecta a uma hidrelétrica,
// ou a uma termoelétrica, ou a energia eólica, ou a uma bateria, que conecta no meu microondas, ou na geladeira, ou
// no aquecedor, ou computador, etc. Porque todos os eletrodomésticos tem um metodo plug de tomada e todas as fontes
// de energia tem a tomada que conecta naquela interface, aquela interface recebe dados das fontes geradoras de
// energia.
// Então, interfaces em Go são a mesma coisa, são maneiras de conectar partes diferentes do nosso código que uma não
// necessariamente sabe o que a outra é, o que ela faz, o que tem por trás, é simplesmente a definição da maneira de
// se conectar.
