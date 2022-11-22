package main

import (
	"fmt"
	"sync"
	"time"
)

const howTime time.Duration = 5

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() { // Читаем из канала
		defer wg.Done()

		for {
			_, ok := <-ch
			if !ok {
				fmt.Println("Закончили чтение ")
				return
			}
		}
	}()

	wg.Add(1)
	go func() { // Пишем в канал
		defer wg.Done()

		for {
			ch <- 1
			select {
			case <-time.After(howTime * time.Second): // Завершение цикла по истечении времени
				close(ch)
				fmt.Println("Закончили запись")
				return
			}
		}
	}()
	wg.Wait() // Ожидание конца работы горутин
}
