package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Задания 6
// Реализовать все возможные способы остановки выполнения горутины.

func main() {
//	1. Самостоятельно. Горутина завершается сама, выполнив возложенные на неё задачи. По окончании
//		она уведомляет о своём завершении.
	var a int 
	a = 1 + a
	return
//	2. Контекст. Сигналом для выхода является истечение переданного контекста.
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// по истечению времени отменяется контекст
	ctx, cancel = context.WithDeadline(ctx, time.Now().Add(time.Second))
	// по истечению времени отменяется контекст
	ctx, cancel = context.WithTimeout(ctx, time.Second)
	fmt.Println(<-ctx.Done())
	defer cancel()
	
//	3. WaitGroup. Воркеры в блоке ждут, когда обнулится переданный sync.WaitGroup.
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hi")
	}()
	wg.Wait()

//	4. Принудительная закрыти с помощью оs
	os.Exit(0)

//	5. C помошью клавишь Ctrl+C, также с получением сигнала для плавного завершения работы
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigint
	fmt.Println("Завершения программы")
	os.Exit(0)
}