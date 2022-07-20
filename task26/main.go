package main

import (
	"fmt"
)

// Задания 26
// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна
// быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func main() {
	fmt.Println(noRepeat("abcd"))
	fmt.Println(noRepeat("abCdefAaf"))
	fmt.Println(noRepeat("aabcd"))
}

func noRepeat(s string) bool {
	for i := range s {
		for j := range s {
			if s[i] == s[j] && i != j {
				return false
			}
		}
	}
	return true
}