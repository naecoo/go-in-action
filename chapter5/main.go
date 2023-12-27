package main

import (
	"fmt"
)

type Person struct {
	name string
	age  uint8
}

func main() {
	fmt.Println("Chapter 5")

	person := &Person{
		name: "oceanpeng",
		age:  25,
	}
	fmt.Println(person)

	var p1 Person
	p1 = Person{"11", 232}
	fmt.Println(p1)

	type Admin struct {
		person Person
		level  string
	}

	a1 := Admin{
		person: Person{"naecoo", 25},
		level:  "123",
	}
	fmt.Println(a1)
}
