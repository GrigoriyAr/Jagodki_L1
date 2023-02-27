/*
Задание:
Реализовать бинарный поиск встроенными методами языка.
*/


package main

import (
	"fmt"
)

func binarySearch(arr []int, elem int) int {
	start := 0 // Начало и конец списка
	end := len(arr) - 1

	for {
		mid := (start + end) / 2 // Делим список на 2 и отсеиваем ненужную половину
		if arr[mid] == elem {    // Если угадали сразу - возвращаем
			return mid
		} else if elem < arr[mid] { // Сдвигаем область поиска влево
			end = mid - 1
		} else { // Сдвигаем область поиска вправо
			start = mid + 1
		}
		if start < 0 || end >= len(arr) { // Если элемента нет в списке возвращаем -1
			return -1
		}
	}
}

func main() {
	arr := []int{-7, -1, 7, 16, 24, 39, 52, 66, 89, 99}
	item := arr[1]
	fmt.Printf("Item = %d, index = %d", item, binarySearch(arr, -1))
}
