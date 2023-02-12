/*
Задание:
Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/


package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	ch := make(chan int)      // Создаем канал
	for _, num := range arr { // Пробегаемся по слайсу
		go func(num int) { // Запускаем горутину с каналом, в который записываем условие
			ch <- num * num
		}(num)
		fmt.Println(<-ch) // Читаем из канала
	}

	time.Sleep(500 * time.Millisecond)
}
