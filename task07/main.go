package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Заданния 7
// Реализовать конкурентную запись данных в map.

func main() {
	var arr = make(map[string]int)
	gorutine(arr)
	channel(arr)
}

func gorutine(arr map[string]int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var worker = 10

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()

			time.Sleep(time.Millisecond)

			mu.Lock()
			arr[randStr(n)] = n * n
			mu.Unlock()
		}(i+1)
	}

	wg.Wait()

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func channel(arr map[string]int) {
	
}


// Генератор рандомных строк A-E 
func randStr(l int) string {
	var pool = "ABCDE"
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}