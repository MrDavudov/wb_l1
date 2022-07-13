package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Задание 4
// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор
// из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения
// работы всех воркеров

func main() {
	var n int
	fmt.Scanf("Введите число воркеров: %d\n", &n)

	// создаём контекст с отменой для плавного завершения
	ctx, cancel := context.WithCancel(context.Background())
	// подписываемся на сигнал остановки от ОС
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	// эта горутина ждёт сигнала от ОС и завершает контекст, давая тем самым сигнал остановить работу
	go func() {
		<-sigint
		fmt.Println("Shutting down...")
		cancel()
	}()
	workerPool(ctx, n)
	fmt.Println("Bye!")
	os.Exit(0)
}

func workerPool(ctx context.Context, w int) {

	// создаем waitgroup для синхронизации
	wg := &sync.WaitGroup{}
	// канал для отправки данных в канал
	numToProcess := make(chan int)

	// количество воркеров которые будут работать
	for i:=0; i<=w; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numToProcess)
		}()
	}

	// отправка данных по каналу
	go func() {
		for {
			var r int
			r = rand.Intn(100)
			time.Sleep(time.Millisecond)
			numToProcess <- r
		}
		// незабываем закрыть канал чтобы не было дед лока
		close(numToProcess)
	}()

	wg.Wait()

}

func worker(ctx context.Context, toProcess <-chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcess:
			// проверка на закрытия канала
			if !ok {
				return
			}
			time.Sleep(time.Millisecond)
			fmt.Printf("Полученое число: %d\n", value)
		}
	}
}
