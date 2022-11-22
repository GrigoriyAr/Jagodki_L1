package main

import "fmt"

// Удалить элемент из слайса можно поменяв его порядок или сохранив его порядок

func quickDelete(sl []string, i int32) {
	sl[i] = sl[len(sl)-1] // Копировать последний элемент в индекс i
	sl[len(sl)-1] = ""    // Удалить последний элемент(записать нулевое значение)
	sl = sl[:len(sl)-1]   // Усечь срез
}

func slowDelete(sl []string, i int32) {
	copy(sl[i:], sl[i+1:]) // Сдвиг [sl + 1:] влево на один индекс
	sl[len(sl)-1] = ""     // Удалить последний элемент(записать нулевое значение)
	sl = sl[:len(sl)-1]    // Усечь срез
}

func main() {
	sl := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Println(sl)
	var i int32 = 2

	quickDelete(sl, i)

	fmt.Println(sl)

	i = 2
	slowDelete(sl, i)
	fmt.Println(sl)

	// Либо можно удалить элемент комбинацией функции append и оператор среза :
	sl = append(sl[:i], sl[i+1])
	fmt.Println(sl)
}
