package main

import "fmt"

func inChan1(num [5]int) <-chan int { // Пишем в канал данные из массива
	out := make(chan int)
	go func() {
		for _, n := range num {
			out <- n
		}
		close(out)
	}()
	return out
}

func inChan2(in <-chan int) <-chan int { // Считываем данные и возвращаем их квадрат
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	num := [5]int{2, 4, 6, 8, 10}
	for square := range inChan2(inChan1(num)) {
		fmt.Println(square)
	}
}
