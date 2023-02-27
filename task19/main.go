/*
Задание:
Разработать программу, которая переворачивает подаваемую на вход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/


package main

import "fmt"

func reverseString(str string) (result string) {
	for _, v := range str { // Пробегаемся циклом по строке, которую принимает ф-ия
		result = string(v) + result // Добавляем каждый символ перед результатом
	}
	return
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(reverseString(str))
}
