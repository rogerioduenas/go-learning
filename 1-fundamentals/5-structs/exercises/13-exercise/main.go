// Create a function that receives a composite structure as a parameter and returns a new structure with all fields modified.
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

func updateStudent(stc student) student {
	stc.name = "Maria"
	stc.age = 50
	stc.course = "Mathematics"
	stc.campus = "LA"
	return stc
}

func main() {
	user := student{
		person: person{"Mike", 18},
		course: "IT",
		campus: "NY",
	}

	fmt.Println(user)
	fmt.Println(updateStudent(user))
}
