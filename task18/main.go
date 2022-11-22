package main

import (
	"fmt"
	"sync"
)

type counter struct {
	i int
	sync.Mutex
}

func (c *counter) increment() { // Увеличивает счетчик
	c.Lock()
	defer c.Unlock()
	c.i++
}

func main() {
	count := new(counter)

	wg := sync.WaitGroup{}

	wg.Add(1000) // Добавляем в waitgroup 1000 задач

	for i := 0; i < 1000; i++ { // Запускаем конкурентные подпрограммы
		go func(c *counter, wg *sync.WaitGroup) {
			defer wg.Done() // Закрываем таску в waitgroup
			c.increment()
		}(count, &wg)
	}
	wg.Wait()
	fmt.Println(count.i)
}
