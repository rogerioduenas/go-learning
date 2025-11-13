// Create a second structure that incorporates the first, representing a compositional relationship between types.
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
	fmt.Println(user.name)
	fmt.Println(user.age)
	fmt.Println(user.course)
	fmt.Println(user.campus)
}
