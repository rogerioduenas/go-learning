// Simulate a type hierarchy where two different structures share the same base structure and create instances for both, showing how inherited fields are reused.
package main

import "fmt"

type person struct {
	name string
}

type employee struct {
	person
	position string
}

type customer struct {
	person
	customerID int
}

func main() {

	myEmployee := employee{
		person:   person{"Mike"},
		position: "Engineer",
	}

	myCustomer := customer{
		person{"Anna"},
		72897398729,
	}

	fmt.Println(myEmployee)
	fmt.Println(myCustomer)
}
