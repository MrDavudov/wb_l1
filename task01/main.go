package main

import "fmt"

// Задание 1
// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).

type HumanAction interface {
	toSpeak()
}

type Human struct {
	Name	string
	Age		int
}

type Action struct {
	Human
	Speak	string
}

func (a *Action) toSpeak(msg string) Action {
	a.Speak = msg
	return *a
}

func main() {
	person := Action{Human{Name: "Rick", Age: 89}, "Hi"}
	fmt.Printf("Type: %T Value: %#v\n", person, person)

	person.toSpeak("Hello")
	fmt.Printf("Type: %T Value: %#v\n", person, person)

	person2 := Action{Human{Name: "Morti", Age: 15}, "Damn"}
	fmt.Printf("Type: %T Value: %#v\n", person2, person2)

	person2.toSpeak("Happy")
	fmt.Printf("Type: %T Value: %#v\n", person2, person2)
}