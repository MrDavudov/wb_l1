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
	
	fmt.Println("---chanel")
	for i, v := range channel(arr) {
		fmt.Println(i, v)
	}
}

func gorutine(arr map[string]int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var w = 10

	for i := 0; i < w; i++ {
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

func channel(arr map[string]int) map[string]int {
	var chInt = make(chan map[string]int)

	go func(chInt chan<- map[string]int) {
		time.Sleep(time.Millisecond)
		v := make(map[string]int)
		for i := 1; i < 10; i++ {
			v[randStr(i)] = i * i
		}
		chInt <- v
	}(chInt)

	return <-chInt
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