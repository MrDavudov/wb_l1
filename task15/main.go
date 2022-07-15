package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// Не хватает функции createHugeString, которая бы возвращала срез из
// 1024 элементов типа string.

// Из-за того, что мы присваиваем переменной justString только 100 значений,
// а не весь срез - 924 элемента болтаются в памяти. При этом, garbage collector
// не может удалить неиспользуемые элементы т.к. срез justString ссылается на
// основной срез.
// Другими словами justString является глобальной переменной, и при выходе с
// функции someFunc() произойдет утечка памяти


func someFunc() {
	var justString string
	v := createHugeString(1 << 10) 
	justString = v[:100]
	fmt.Println(justString)
}

func createHugeString(n int) string {
	str := strings.Repeat("X", n)
	return str
}

func main() {
  someFunc()
}