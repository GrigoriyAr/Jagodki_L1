/*
Задание:
Реализовать пересечение двух неупорядоченных множеств.
*/


package main

import "fmt"

func intersect(arr1, arr2 *[]int) []int {
	var result []int
	set := make(map[int]struct{})

	if len(*arr1) >= len(*arr2) { // Проверяем, чтобы массив, с которым будут сравнивать, был длиннее или такой же
		for _, v := range *arr1 { // Записываем в мапу из первого массива
			set[v] = struct{}{}
		}
		for _, v := range *arr2 { // Пробегаемся по второму массиву и записываем совпадения в result
			_, ok := set[v]
			if ok {
				result = append(result, v)
			}
		}
		return result
	} else { // Если длиннее второй массив
		for _, v := range *arr2 { // Записываем в мапу из второго массива
			set[v] = struct{}{}
		}
		for _, v := range *arr1 { // Пробегаемся по первому массиву и записываем совпадения в result
			_, ok := set[v]
			if ok {
				result = append(result, v)
			}
		}
		return result
	}
}
func main() {
	arr1 := []int{1, 2, 13, 26, 19, 6, 9, 67, 4, 24}
	arr2 := []int{18, 11, 2, 6, 19, 51, 24, 25, 5, 3, 1}

	fmt.Println("Первое множество: ", arr1)
	fmt.Println("Второе множество: ", arr2)
	intersectRes := intersect(&arr1, &arr2)
	fmt.Println("Точки пересечений: ", intersectRes)
}
