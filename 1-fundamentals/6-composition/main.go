package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type student struct {
	person
	course string
	campus string
}

func main() {

	person := person{name: "Mike", age: 21}
	student := student{person, "Computer Science", "London"}

	fmt.Println(student)
	fmt.Println(student.name)
	fmt.Println(student.course)
}
