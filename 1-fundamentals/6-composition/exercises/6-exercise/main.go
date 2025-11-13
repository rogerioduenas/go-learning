// Access a field inherited from the base structure directly through the composite structure and display its value.
package main

import "fmt"

type person struct {
	name string
}

type student struct {
	person
}

func main() {

	user := student{
		person: person{
			name: "Mike",
		},
	}

	fmt.Println(user.name)
	//or
	fmt.Println(user.person.name)
}
