/*
Задание:
Разработать программу, которая перемножает, делит,
складывает, вычитает две числовых переменных a, b, значение которых > 2^20.
*/


package main

import (
	"fmt"
	"math/big"
)

// В данной задаче используем пакет math/big для работы с длинной арифметикой
func main() {
	a, _ := new(big.Int).SetString("222333444555666777888999", 10) // Инициализируем значение из строки
	b, _ := new(big.Int).SetString("444445555566660003", 10)       //Методом SetString

	// сложение
	sum := a.Add(a, b)
	fmt.Printf("a + b = %s\n", sum)

	// вычитание
	sub := a.Sub(a, b)
	fmt.Printf("a - b = %s\n", sub)

	// деление
	div := a.Div(a, b)
	fmt.Printf("a / b = %s\n", div)

	// умножение
	mult := a.Mul(a, b)
	fmt.Printf("a * b = %s\n", mult)
}
