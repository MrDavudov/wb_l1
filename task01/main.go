package main

import "fmt"

// Задание 1
// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

type Human struct {
	Name	string
	Age		int
}

func (n *Human) newName() Human {
	n.Name = "Devid"
	return n
}

func One() {
	human := Human{
		"Jack",
		12,
	}
	newName(human)
	fmt.Printf("New name: %s", human.Name)
}