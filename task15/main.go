package main

import (
	"fmt"
	"strings"
)

/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

Происходит утечка памяти. В первую очередь, использование глобальных переменных не желательно, по этому можно либо записать ее внутри ф-ии,
либо передать в функцию по указателю. У justString фиксированная длина, использование локальной переменной
было бы лучше. После создания слайса 924 элемента не будут использоваться и лишняя память не будет очищаться.
*/
// Корректная реализация:

func someFunc(justString *string) {
	v := createHugeString(1 << 10)
	*justString = strings.Clone(v)[:100] // Выделяем новый участок памяти для скопированной строки
}
func createHugeString(len int) string {
	res := ""
	for i := 0; i < len; i++ {
		res += "1"
	}
	return res
}
func main() {
	var justString string
	someFunc(&justString)
	fmt.Println(justString)
}
