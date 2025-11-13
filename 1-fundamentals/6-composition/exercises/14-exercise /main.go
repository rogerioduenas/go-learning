// Create a function that receives the derived structure and modifies both inherited fields and specific fields before returning it.
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

func modifyWorkplaceAndAge(emp staff) staff {
	emp.workplace = "NY"
	emp.age = 22
	//or
	// emp.person.age = 22
	return emp
}

func main() {

	var employee staff = staff{person: person{"Mike", 21}, position: "Engineer", workplace: "LA"}

	fmt.Println(employee)
	fmt.Println(modifyWorkplaceAndAge(employee))
}
