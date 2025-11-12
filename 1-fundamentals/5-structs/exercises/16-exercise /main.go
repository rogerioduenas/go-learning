// Create two related structures, one composed of the other. Create an instance using partial field assignment and then replace some internal values ​​to demonstrate data updating.
package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type employee struct {
	person
	salary   float64
	discount uint
}

func main() {
	emp := employee{
		person: person{
			name: "Mike",
		},
		salary: 2000,
	}
	fmt.Println(emp)

	emp.name = "Anna"
	emp.salary = 3000
	fmt.Println(emp)
}
