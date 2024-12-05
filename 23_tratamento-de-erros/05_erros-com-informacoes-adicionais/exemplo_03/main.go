package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("norgate math again: square root of negative number: %v", x)
	}
	return math.Sqrt(x), nil
}
