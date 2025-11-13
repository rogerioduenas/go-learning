// Display the value of a field belonging to the innermost structure using the outermost structure.
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

	myEmployee := salary{
		employee: employee{
			person:   person{"Anna"},
			position: "Architect",
		},
		tc: 3000,
	}
	fmt.Println(myEmployee.name)
	fmt.Println(myEmployee.employee.person.name)

}
