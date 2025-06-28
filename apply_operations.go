package main

import (
	"fmt"
)

func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func main() {
	a, b := 10, 5

	sum := applyOperation(a, b, add)
	diff := applyOperation(a, b, subtract)
	prod := applyOperation(a, b, multiply)

	fmt.Println("Сложение:", sum)
	fmt.Println("Вычитание:", diff)
	fmt.Println("Умножение:", prod)
}
