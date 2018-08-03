package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Rajib", 38})

	fmt.Println(person{name: "Zaman", age: 36})

	fmt.Println(person{name: "Jimmy"})

	fmt.Println(&person{name: "Ann", age: 40})

	aperson := person{name: "Hulk", age: 36}
	fmt.Println(aperson.name, aperson.age)

	anotherperson := &aperson
	fmt.Println(anotherperson.name, anotherperson.age)

	anotherperson.age = 35
	fmt.Println(aperson.age)
}
