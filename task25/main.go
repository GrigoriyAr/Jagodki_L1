/*Задание:
Реализовать собственную функцию sleep.
*/


package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	SleepAfter(time.Second * 3)
	sleepTime(time.Second * 3)
	sleepContext(time.Second * 3)
}

func SleepAfter(tm time.Duration) { // Спит пока After не попадет в канал
	fmt.Printf("Спим %.2f секунды\n", tm.Seconds())
	if <-time.After(tm); true {
		return
	}
}

func sleepTime(tm time.Duration) { // Спим пока не завершится таймер
	fmt.Printf("Спим %.2f секунды\n", tm.Seconds())
	until := time.After(tm)
	select {
	case <-until:
		return
	}
}

func sleepContext(tm time.Duration) { // Ждет завершения контекста
	fmt.Printf("Спим %.2f секунды\n", tm.Seconds())
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, tm)
	defer cancel()

	if <-ctx.Done(); true {
		return
	}
}
