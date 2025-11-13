// Implement a chain of structures, where a base structure is incorporated into an intermediate one, which in turn is incorporated into a third, representing multiple levels of composition.
package main

import "fmt"

type person struct {
	name string
}

type employee struct {
	person
	position string
}

type salary struct {
	employee
	tc float64
}

func main() {

	var myEmployee salary
	myEmployee.name = "Mike"
	myEmployee.position = "Engineer"
	myEmployee.tc = 2000
	fmt.Println(myEmployee)

	//or

	myEmployee2 := salary{
		employee: employee{
			person:   person{"Anna"},
			position: "Architect",
		},
		tc: 3000,
	}
	fmt.Println(myEmployee2)

}
