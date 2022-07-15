package main

// Задания 19
// Разработать программу, которая переворачивает подаваемую на ход строку 
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

import "fmt"

func main() {
	str := "главрыба"
	fmt.Println(str)

	res1 := ""
	for _, v := range str {
		res1 = string(v) + res1
	}

	fmt.Println(res1)
}