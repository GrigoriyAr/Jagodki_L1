/*
Задание:
Реализовать конкурентную запись данных в map.
*/


package main

import (
	"errors"
	"log"
	"sync"
)

// Гонки данных происходят, когда несколько горутин одновременно обращаются к одному и тому же ресурсу
// и по крайней мере один из обращений является записью

type safeNumbers struct { // Встраиваем мьютекс в структуру
	sync.RWMutex
	numbers map[int]int
}

func (sn *safeNumbers) Get(num int) (int, error) {
	sn.RLock()
	defer sn.RUnlock()
	if number, ok := sn.numbers[num]; ok {
		return number, nil
	}
	return 0, errors.New("Номер не существует")
}

func (sn *safeNumbers) Add(num int) {
	sn.Lock()
	defer sn.Unlock()
	sn.numbers[num] = num
}

func generateNumbersMap() {
	wg := sync.WaitGroup{}
	safeNumber := &safeNumbers{ // Инициализируем безопасную структуру
		numbers: map[int]int{},
	}
	for i := 0; i < 100; i++ { // Записываем
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeNumber.Add(i)
		}(i)
	}
	for i := 0; i < 100; i++ { // Читаем
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			number, err := safeNumber.Get(i)
			if err != nil {
				log.Print(err)
			} else {
				log.Print(number)
			}
		}(i)
	}

	wg.Wait()
}

func main() {
	generateNumbersMap()
}
