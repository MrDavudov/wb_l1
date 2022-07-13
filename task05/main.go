package main

import (
	"fmt"
	"sync"
	"time"
)

// Задание 5
// Разработать программу, которая будет последовательно отправлять 
// значения в канал, а с другой стороны канала — читать. По истечению 
// N секунд программа должна завершаться.

func main() {
	var count int
	var wg sync.WaitGroup
	var mu sync.RWMutex

	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// если в момент использовании функции, нет запись происходит чтение
			time.Sleep(time.Nanosecond)

			mu.RLock()
			_ = count
			mu.RUnlock()
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()

			// экслючивная блокировка пока не завершится запись
			time.Sleep(time.Nanosecond)

			mu.Lock()
			count++	
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(count)
	fmt.Printf("Время на выполения: %v", time.Since(start).Seconds())

}