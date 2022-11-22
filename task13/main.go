package main

import "fmt"

func wayOne(a, b int) { // свап методом присвоения
	b, a = a, b
	fmt.Println(a, b)
}

func wayTwo(a, b int) { // Арифметический свап
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a, b)
}

func wayThree(a, b int) { // Свап через побитовые операции ( при помощи исключающего ИЛИ)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}

func main() {
	a, b := 3, 5
	wayOne(a, b)
	wayTwo(a, b)
	wayThree(a, b)
}
