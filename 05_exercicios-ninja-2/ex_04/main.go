package main

import "fmt"

func main() {
	x := 200
	fmt.Printf("%d, %b, %#x", x, x, x)

	y := x << 1
	fmt.Printf("%d, %b, %#x", y, y, y)
}
