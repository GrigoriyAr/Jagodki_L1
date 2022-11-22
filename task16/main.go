package main

import "fmt"

func partition(arr []int, low, high int) ([]int, int) { // Поиск опорного элемента и перемещение элементов относительно его
	pivot := arr[high] // Опорный элемент, выбираем самый правый элемент массива
	i := low
	for j := low; j < high; j++ { //Разделяем слайс на две части - первый и последний элемент
		if arr[j] < pivot { //  Все, что больше, окажется справа от i
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i] // Элемент, относительно которого мы распределяли, ставим на позицию i
	return arr, i
}
func quickSort(arr []int, low, high int) []int { // Продолжаем сортировку, пока левая и правая границы слайса не совпадут
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}
func main() {
	fmt.Println(quickSortStart([]int{5, 16, 72, 23, 10, 0, 14, 99}))
}
