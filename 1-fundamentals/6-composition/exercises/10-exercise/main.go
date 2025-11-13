// Create a new instance of the derived structure by initializing all fields directly, without creating an intermediate variable for the base structure.
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
}
