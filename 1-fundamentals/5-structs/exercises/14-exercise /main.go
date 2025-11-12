// Develop a function that creates and directly returns an instance of a structure, populating its fields internally.
package main

import "fmt"

type person struct {
	name string
	age  uint8
}

func createUser(name string, age uint8) person {
	result := person{
		name: name,
		age:  age,
	}
	return result
}

func main() {
	fmt.Println(createUser("Mike", 19))
	fmt.Println(createUser("Anna", 21))
}
