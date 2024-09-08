package main

import "fmt"

func main() {
	a := (10 == 10)
	b := (10 <= 100)
	c := (10 < 100)
	d := (10 >= 100)
	e := (10 > 100)
	f := (10 != 100)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
