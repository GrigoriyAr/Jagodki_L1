package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func worker(i int, ch <-chan int) {
	for data := range ch { // Пишем число в поток
		time.Sleep(time.Millisecond * 100)
		fmt.Fprintf(os.Stdout, "Worker: %d    | Value recieved: %d\n", i, data)
	}
}

// После завершения main воркеры завершаются

func main() {
	var wCount int // Выбираем кол-во воркеров
	fmt.Scanf("%d", &wCount)
	ch := make(chan int) // Передаем данные
	for i := 0; i < wCount; i++ {
		go worker(i+1, ch)
	}

	done := make(chan os.Signal, 1)                                    // Определяем канал для обработки прерывания
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // Завершаем процесс ctrl c

	for { // Запускаем бесконечный цикл
		select {
		case <-done:
			close(ch)
			fmt.Fprintln(os.Stdout, "Program exited, interrupted")
			return
		default:
			ch <- rand.Intn(100)
		}
	}
}
