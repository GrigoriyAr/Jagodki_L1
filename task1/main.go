package main

import "fmt"

type Human struct {
	name  string
	age   int
	adult bool
}

type Action struct { // Структура, которая хранит значения Human, у которой еще есть свойство, которое хранит работу
	*Human
	work string
}

func main() {
	person := &Action{ //Мы можем работать с инстансом объекта Action так,
		&Human{ //будто у него есть набор свойств структуры  Human
			name:  "Ivan",
			age:   25,
			adult: true,
		},
		"developer",
	}
	fmt.Printf("%+v\n", person)
	fmt.Printf("Меня зовут %v и я работаю %v", person.name, person.work)
}
