package main

import (
	"fmt"
	"sync"
)

// Задания 18
// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	Count	int
}

func main() {
	count := Counter{}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()
			count.Count++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}