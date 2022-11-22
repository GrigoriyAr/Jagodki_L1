package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func UsingChanClose(ch <-chan int, wg *sync.WaitGroup) {
	for i := range ch {  // Читаем данные из канала, пока он не будет закрыт
		fmt.Print(i, " ")
	}
	fmt.Print("Горутина остановилась\n")
	wg.Done()
	return
}

func UsingChan(exit <-chan interface{}, wg *sync.WaitGroup) {
	for {     // выводим данные, пока в канал выхода не будет совершена запись
		select {
		case <-exit:
			fmt.Print("Горутина остановилась\n")
			wg.Done()
			return
		default:
			fmt.Print("Горутина запустилась\n")
			time.Sleep(1 * time.Second)
		}
	}
}

func UsingContext(ctx context.Context, wg *sync.WaitGroup) {
	for {   // Выводим данные, пока контекст не будет закрыт
		select {
		case <-ctx.Done():
			fmt.Print("Горутина остановилась\n")
			wg.Done()
			return
		default:
			fmt.Print("Горутина запустилась\n")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup   // остановка выполнения горутины с использованием закрытия канала
	wg.Add(1)
	ch := make(chan int)
	go UsingChanClose(ch, &wg)
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()

	wg.Add(1)    // остановка выполнения горутины с использованием канала закрытия
	exit := make(chan interface{})
	go UsingChan(exit, &wg)
	time.Sleep(1 * time.Second)
	exit <- struct{}{}
	wg.Wait()
	
	ctx, cancel := context.WithCancel(context.Background())  // остановка выполнения горутины с использованием контекстов
	wg.Add(1)
	go UsingContext(ctx, &wg)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}
