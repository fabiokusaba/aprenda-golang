package main

import "fmt"

type erroEspecial struct {
	message string
}

func (e erroEspecial) Error() string {
	return fmt.Sprintf("vixe, deu erro cara! %v", e.message)
}

func recebeErr(e error) {
	fmt.Println(e)
}

func main() {
	err := erroEspecial{"fuuuuuuuuu..."}
	recebeErr(err)
}
