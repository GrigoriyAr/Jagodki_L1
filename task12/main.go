package main

import "fmt"

func MakeSet(arr []string) []string {
	data := make(map[string]bool) // Храним множество в мапе
	for _, i := range arr {
		data[i] = true
	}
	resSet := make([]string, 0, len(data))
	for key := range data {
		resSet = append(resSet, key)
	}
	return resSet // Возвращаем слайс уникальных строк
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(MakeSet(arr))
}
