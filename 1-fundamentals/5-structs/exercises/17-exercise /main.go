// Create a function that receives multiple structures, combines their fields into a new instance, and returns the result.
package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type employee struct {
	salary   float64
	discount uint
}

type combined struct {
	person
	employee
}

func joinStructs(p person, e employee) combined {

	return combined{
		person:   p,
		employee: e,
	}
}

func main() {
	fmt.Println(joinStructs(person{name: "Mike", age: 21},
		employee{salary: 2000, discount: 10}))
}
