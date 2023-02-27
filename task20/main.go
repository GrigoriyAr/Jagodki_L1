/*
Задание:
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(str string) string {
	strArray := strings.Split(str, " ") // Получаем слайс из исходной строки по пробелу
	resString := ""
	for i := len(strArray) - 1; i >= 1; i-- { // Пробегаем в обратном порядке
		resString += strArray[i] + " " // Складываем с результирующей строкой
	}
	resString += strArray[0]
	return resString
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Printf("В обратном порядке: %s", reverse(s))
}
