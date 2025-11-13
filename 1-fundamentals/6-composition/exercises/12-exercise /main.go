// Create a function that receives a derived structure and displays a field inherited from the base structure.
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

func showBaseData(stc staff) string {
	return stc.name
}

func main() {
	myEmployee := staff{
		person:    person{"Mike", 21},
		position:  "Engineer",
		workplace: "LA",
	}

	fmt.Println(showBaseData(myEmployee))
}
