package main

import (
	"fmt"
	"strings"
)

func checkUnique(str string) bool {
	str = strings.ToLower(str) // Переводим все символы строки в нижний регистр
	m := make(map[rune]int)
	for _, ch := range str { //Пробегаемся по символам строки
		if _, ok := m[ch]; !ok { // Ключ с текущим символом не нашелся
			m[ch] = 0
		} else {
			// Иначе, если ключ нашелся, то символ уже был и нужно вернуть false
			return false
		}
	}
	return true // Все символы уникальные
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(checkUnique(str))
}
