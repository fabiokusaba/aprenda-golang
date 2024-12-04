package main

import "fmt"

func main() {
	c := make(chan int)
	go botala(c)
	
	for v := range c {
		fmt.Println(v)
	}
}

func botala(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
