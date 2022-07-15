package main

import (
	"fmt"
	"strings"
)

// Задания 20
// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func main() {
	str := "snow dog sun"
	fmt.Println(str)

	var res string
	var s string
	for i, v := range str {
		s = s + string(v)
		if string(v) == " " || len(str)-1 == i {
			res = strings.Trim(s, " ") + " " + res
			s = ""
		}
	}

	fmt.Println(res)
}