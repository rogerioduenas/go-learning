// Display both inherited and derived structure-specific fields in the standard output.
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

	worker := person{"Mike", 21}

	myEmployee := staff{
		person:    worker,
		position:  "Engineer",
		workplace: "LA",
	}

	fmt.Println(myEmployee)
}
