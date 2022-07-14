package main

import "fmt"

// Задания 12
// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

type animal struct{}

type set map[string]animal

var member animal

func main() {
	var str = []string{"cat", "cat", "dog", "cat", "tree"}

	animals := set{}

	for _, v := range str {
		animals[v] = member
	}

	fmt.Println(animals)
}