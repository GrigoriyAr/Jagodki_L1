package main

import "fmt"

func main() {
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float32) // Инициализируем мапу, для которой ключ - шаг, значение - слайс подходящих значений
	for _, v := range temp {
		ten := int(v/10) * 10
		m[ten] = append(m[ten], v)
	}
	fmt.Println(m)
}
