package main

import "fmt"

/*
В данной задаче реализуется паттерн адаптер на примере на примере зарядки для Iphone (которая нам нужна),
но рядом у нас есть только зарядка на android.
Для того, чтобы это зарядить телефон, была реализована структура androidToIphoneAdapter, которая хранит в себе
указатель на зарядку для андроида и выполняет функцию преобразования одного интерфейса взаимодействия в нужный.
*/

type androidPlug struct { // iphonePlug - у нас есть зарядка на айфон
}

func (plug androidPlug) insertPlugToSocket() { // Вставляем в розетку
	fmt.Printf("Вставляем андротдовскую зарядку в розетку.\n")
}

type iphonePlug struct { // Нужная зарядка
}

func (plug iphonePlug) insertPlugToSocket() {
	panic("implement me")
}

func (plug iphonePlug) insertIphonePlugToSocket() { // Так же может вставляться в розетку
	fmt.Printf("Вставляем айфоновский штепсель в розетку.\n")
}

type insertIphonePlug interface { // Подходит только айфоновская, но хочется зарядить андроидовской
	insertPlugToSocket()
	insertIphonePlugToSocket()
}

type androidToIphoneAdapter struct { // Нашли переходник
	Adapter *androidPlug
}

func (adapter androidToIphoneAdapter) insertPlugToSocket() {
	panic("implement me")
}

func (adapter androidToIphoneAdapter) insertIphonePlugToSocket() {
	fmt.Printf("Вставляем нужную зарядку с помощью адаптера.\n")
}

type Client struct {
}

func (client Client) ChargePhone(plug insertIphonePlug) { // Пользуемся нужной зарядкой
	plug.insertIphonePlugToSocket()
}

func main() {
	client := Client{}
	iphonePlugExemplar := iphonePlug{}
	client.ChargePhone(iphonePlugExemplar)

	euPlugExemplar := androidPlug{}
	adapter := androidToIphoneAdapter{Adapter: &euPlugExemplar}
	client.ChargePhone(adapter)
}
