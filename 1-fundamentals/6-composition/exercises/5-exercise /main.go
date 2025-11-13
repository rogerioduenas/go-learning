// Display the full value of a compound structure in standard output.
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
		course: "IT",
		campus: "LA",
	}

	fmt.Println(user)
}
