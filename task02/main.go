package main

import (
	"fmt"
	"sync"
)

// Задание 2
// Написать программу, которая конкурентно рассчитает значение
// квадратов чисел взятых из массива (2,4,6,8,10) и выведет их
// квадраты в stdout.

// -> конкуретность
// https://medium.com/nuances-of-programming/%D0%BA%D0%BE%D0%BD%D0%BA%D1%83%D1%80%D0%B5%D0%BD%D1%82%D0%BD%D0%BE%D1%81%D1%82%D1%8C-%D0%B8-%D0%BF%D0%B0%D1%80%D0%B0%D0%BB%D0%BB%D0%B5%D0%BB%D0%B8%D0%B7%D0%BC-%D0%B2-golang-go-%D0%BF%D1%80%D0%BE%D1%86%D0%B5%D0%B4%D1%83%D1%80%D1%8B-82bae0f92e81
// https://golangify.com/concurency
// -> что такое sync
//
// -> что такое stdout
// https://metanit.com/go/tutorial/8.6.php
// -> что такое data race
// https://yourbasic.org/golang/data-races-explained/#:~:text=A%20data%20race%20happens%20when,and%20it's%20behavior%20is%20undefined.

func main() {
	var arr = []int{2,4,6,8,10}
	var res int

	// используем sync.WaitGroup, чтобы дождаться окончания работы воркеров.
	wg := new(sync.WaitGroup)
	var mu sync.Mutex

	for _, v := range arr {

		wg.Add(1)
		// передаём число в качестве аргумента горутине, чтобы избежать data race.
		go func(n int) {
			defer wg.Done()

			mu.Lock()
			res = res + n * n
			mu.Unlock()
			
		} (v)
	}
	// ждём окончания
	wg.Wait()
	
	fmt.Printf("Task 2: %v\n", res)
}