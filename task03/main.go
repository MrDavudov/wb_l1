package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Задание 3
// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	var arr = []int{2,4,6,8,10}
	fmt.Printf("Sync method: %v, Chan method: %v\n", syncInt(arr), chanInt(arr, 5))


	const sliceSize = 5000000
	const chanSize = 5000000
	numbers := make([]int, 0, sliceSize)
	for i := 1; i <= sliceSize; i++ {
		numbers = append(numbers, i*2)
	}

	start := time.Now()
	syncMutexRes := syncMuInt(numbers)
	t1 := time.Since(start)
	fmt.Printf("useSyncMutex -> result: %d, time: %v\n", syncMutexRes, t1)
	start = time.Now()
	atomicRes := syncInt(numbers)
	t2 := time.Since(start)
	fmt.Printf("useAtomic -> result: %d, time: %v\n", atomicRes, t2)
	start = time.Now()
	chanRes := chanInt(numbers, chanSize)
	t3 := time.Since(start)
	fmt.Printf("useChannels -> result: %d, time: %v\n", chanRes, t3)
}

func syncMuInt(arr []int) (res int) {
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

	return res
}

func syncInt(arr []int) (res int64) {
	// используем sync.WaitGroup, чтобы дождаться окончания работы воркеров.
	wg := new(sync.WaitGroup)

	for _, v := range arr {
		wg.Add(1)
		// передаём число в качестве аргумента горутине, чтобы избежать data race.
		go func(n int) {
			defer wg.Done()
			// atomic гарантирует, что не будет гонки данных
			atomic.AddInt64(&res, int64(n*n))
		} (v)
	}
	// ждём окончания
	wg.Wait()

	return res
}

func chanInt(arr []int, chSize int) int {
	wg := new(sync.WaitGroup)

	chSqr := make(chan int, chSize) // канал для передачи квадратов
	chResult := make(chan int)      // канал для результата

	for _, v := range arr {
		wg.Add(1)
		go func(n int) {
			// передаем квадрат в канал
			chSqr <- n * n
			wg.Done()
		}(v)
	}
	// эта горутина собирает квадраты из канала и складывает их
	go func() {
		var result int
		// цикл завершится, когда канал будет закрыт
		for v := range chSqr {
			result += v
		}
		chResult <- result
	}()
	// ждём окончания
	wg.Wait()
	// закрытие канала даёт сигнал горутине послать результат
	close(chSqr)
	return <-chResult
}

