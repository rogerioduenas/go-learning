// Create a base structure called person and another composite structure called student that contains person as a field. At the program's entry point, declare a variable for student and assign it the previously created instance of person.
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

	user := student{
		person: person{
			name: "Mike",
			age:  18,
		},
	}

	fmt.Println(user)
	fmt.Println(user.name)
	fmt.Println(user.age)
}
