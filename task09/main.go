package main

import (
	"fmt"
)

// Задания 9
// Разработать конвейер чисел. Даны два канала: в первый пишутся
// числа (x) из массива, во второй — результат операции x*2, после
// чего данные из второго канала должны выводиться в stdout.

func main() {
	var ch1 = make(chan int, 1)
	var ch2 = make(chan int, 1)
	var arr = []int{1, 2, 3, 4, 5}
	
	go func() {
		for _, v := range arr {
			ch1 <- v
		}
		close(ch1)
	}()

	go func() {
		for v := range ch1 {
			ch2 <- 2*v
		}
		close(ch2)
	}()

	for v := range ch2 {
		fmt.Println(v)
	}
}