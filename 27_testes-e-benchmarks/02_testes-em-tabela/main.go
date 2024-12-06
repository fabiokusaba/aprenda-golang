package main

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v // total = total + v
	}
	return total
}

func Multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v // total = total * v
	}
	return total
}
