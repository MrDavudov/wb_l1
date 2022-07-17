package main

// Задания 23
// Удалить i-ый элемент из слайса.

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	i := 0
	// 1. с помощью функции copy
	copy(arr[i:], arr[i+1:])
	fmt.Println(arr)
	// 2. урезать слайс, удалив последний элемент
	arr = arr[:len(arr)-1]
	fmt.Println(arr)
}
