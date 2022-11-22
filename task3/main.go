package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum int64
	num := []int64{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _, v := range num { // Пробегаемся по массиву
		wg.Add(1) // Добавляем кол-во задач в waitgroup
		go func(v int64) {
			defer wg.Done() // Закрываем таску в горутине
			atomic.AddInt64(&sum, v*v)
		}(v)
	}
	wg.Wait() // Блокируем поток пока счетчик не обнулится
	fmt.Printf("сумма квадратов: %d", sum)
}
