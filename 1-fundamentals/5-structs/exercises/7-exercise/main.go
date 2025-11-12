// Create an instance of a composite structure using direct initialization, filling all fields in sequence.
package main

import "fmt"

type student struct {
	course string
	campus string
}

type person struct {
	name string
	age  uint8
	student
}

func main() {
	user := person{
		name: "Mike",
		age:  18,
		student: student{
			course: "IT",
			campus: "NY"}}
	//or
	// user := person{"Mike", 18, student{"IT", "NY"}}
	fmt.Println(user)
}
