package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Задания 16
// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < 1000; i++ {
		arr = append(arr, rand.Intn(100))
	}

	fmt.Println(arr)
	
	quicksort(arr)
	start := time.Now()
	fmt.Printf("arr: %v\ntime: %v", arr, time.Since(start).String())
}

func quicksort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
     
    left, right := 0, len(arr)-1
     
    pivot := rand.Int() % len(arr)
     
    arr[pivot], arr[right] = arr[right], arr[pivot]
     
    for i := range arr {
        if arr[i] < arr[right] {
            arr[left], arr[i] = arr[i], arr[left]
            left++
        }
    }
     
    arr[left], arr[right] = arr[right], arr[left]
     
    quicksort(arr[:left])
    quicksort(arr[left+1:])
     
    return arr
}