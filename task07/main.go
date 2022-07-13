package main

import (
	"fmt"
	"sync"
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
	var worker = 100
	for i := 0; i < worker; i++ {
		wg.Add(100)
		go func(n int) {
			mu.Lock()
			arr[string(n)] = n * n
			mu.Unlock()
		}(i)
	}
	fmt.Println(arr)
}

func channel(arr map[string]int) {
	
}