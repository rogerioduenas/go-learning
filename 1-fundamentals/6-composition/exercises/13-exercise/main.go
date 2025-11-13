// Develop a function that receives the base structure and returns a derived structure containing that embedded instance along with additional fields.
package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type staff struct {
	person
	position  string
	workplace string
}

func createStaffFromPerson(base person) staff {
	employee := staff{
		person:    base,
		position:  "Engineer",
		workplace: "LA",
	}
	return employee
}

func main() {
	base := person{"Mike", 21}

	fmt.Println(createStaffFromPerson(base))
}
