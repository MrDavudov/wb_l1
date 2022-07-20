package main

import (
	"fmt"
	"time"
)

// Задание 25
// Реализовать собственную функцию sleep.

func main() {
	start := time.Now()
	sleep(time.Second*2)

	fmt.Printf("Прошло %s секунд", time.Since(start).String())
}

func sleep(d time.Duration) {
	<-time.After(d)
}

