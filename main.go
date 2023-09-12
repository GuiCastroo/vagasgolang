package main

import "fmt"

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	p := Person{name: name, age: age}
	return &p
}

func main() {
	joao := NewPerson("João", 20)
	pedro := NewPerson("Pedro", 10)
	fmt.Println(joao, pedro)
}
