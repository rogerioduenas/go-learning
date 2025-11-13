// Create a base structure representing a generic type of worker and another derived structure that adds specific fields, such as job title and workplace.
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

func main() {

	employee := staff{
		person: person{
			name: "Mike",
			age:  18,
		},
		position:  "Engineer",
		workplace: "LA",
	}

	fmt.Println(employee)
}
