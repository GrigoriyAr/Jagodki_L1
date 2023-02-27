/*
Задание:
Дана переменная int64. Разработать программу, которая устанавливает i-й бит в 1 или 0
*/

package main

import "fmt"

func main() {
	var num int64 // Число
	var i int64   // Индекс
	fmt.Scan(&num)
	fmt.Scan(&i)
	fmt.Printf(" %d в двоичной форме будет: %b\n", num, num)
	inv := 1 << i // // Инициализация числа, i-ый бит которого противоположен i-ому биту числа num
	fmt.Printf("%d\n в двоичной форме будет: %b\n", num^int64(inv), num^int64(inv))
}
