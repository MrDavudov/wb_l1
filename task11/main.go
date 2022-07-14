package main

// Задания 11
// Реализовать пересечение двух неупорядоченных множеств.

import "fmt"

type void struct{}

/* Объявляем тип set, у которого вместо значения пустая структура(пустая структура не занимает место в памяти)
Передав этот тип переменной, мы автоматически делаем ее множеством*/
type set map[int]void

// Объявляем общую переменную с пустой структурой, которую мы будем использовать как значение для множества
var member void

func intersect(a, b []int) set {

	out := set{}

	for _, x := range a {
		for _, y := range b {
			if x==y{
				out[x] = member
				break
			}
		}
	}
	
	return out
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{7, 15, 3, 22, 5, 1, 9}

	fmt.Println(intersect(s1, s2))
}