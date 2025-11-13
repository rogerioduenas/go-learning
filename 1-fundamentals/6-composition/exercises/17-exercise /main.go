// Create a function that receives two employee-type structures, uses the inherited data from each—such as name and job title—and combines this information to generate a new salary-type structure, incorporating the inherited fields and adding a salary value.
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

func joinStructs(emp1 employee, emp2 employee) salary {

	return salary{
		employee: employee{
			person:   emp1.person,
			position: emp2.position,
		},
		tc: 2500,
	}
}
func main() {

	name := employee{person: person{"Mike"}}
	position := employee{position: "Engineer"}

	fmt.Println(joinStructs(name, position))

}
