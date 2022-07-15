package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Задания 17
// Реализовать бинарный поиск встроенными методами языка.

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(100))
	}

	x := 5
	count := 0

	for _, v := range arr {
		if x == v {
			count++
		}
	}
	fmt.Printf("найден: %d, в количестве %d", x, count)
}