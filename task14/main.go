/*
Задание:
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/


package main

import "fmt" //"reflect"

func variableType(variable interface{}) {
	switch variable.(type) {
	case int:
		fmt.Println("Переменная типа int")
	case string:
		fmt.Println("Переменная типа string")
	case bool:
		fmt.Println("Переменная типа bool")
	case chan int:
		fmt.Println("Переменная типа channel")
	case chan string:
		fmt.Println("Переменная типа channel")
	case chan bool:
		fmt.Println("Переменная типа channel")

	}
}

// Либо можно использовать пакет reflect
// Берем значение типа interface{} и извлекаем информацию о его типе вызовом TypeOf
/* func variableType(variable interface{}) {
	switch reflect.TypeOf(variable).Kind() {
	case reflect.Int:
		fmt.Println("Переменная типа int")
	case reflect.String:
		fmt.Println("Переменная типа string")
	case reflect.Bool:
		fmt.Println("Переменная типа bool")
	case reflect.Chan:
		fmt.Println("Переменная типа channel")
	}
}
*/

func main() {
	var a interface{} = 1
	var b interface{} = "1"
	var c interface{} = true
	var d interface{} = make(chan int)
	var e interface{} = make(chan string)
	var f interface{} = make(chan bool)

	variableType(a)
	variableType(b)
	variableType(c)
	variableType(d)
	variableType(e)
	variableType(f)

}
