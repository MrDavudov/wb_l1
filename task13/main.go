package main

import "fmt"

// Задание 13
// Поменять местами два числа без создания временной переменной.

func main() {
	a := 1
	b := 2
	fmt.Printf("a: %v, b: %v\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a: %v, b: %v\n", a, b)
}