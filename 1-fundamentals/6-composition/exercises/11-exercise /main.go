// Change the value of an inherited field after the derived structure has been created and display the updated result.
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
	myEmployee := staff{
		person:    person{"Mike", 21},
		position:  "Engineer",
		workplace: "LA",
	}

	fmt.Println(myEmployee)

	myEmployee.name = "Anna"

	fmt.Println(myEmployee)
}
