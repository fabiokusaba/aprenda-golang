package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNorgateMath
	}
	return math.Sqrt(x), nil
}
