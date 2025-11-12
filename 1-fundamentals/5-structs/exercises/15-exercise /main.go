// Create a composite type called Person that contains fields of different types, including a field based on another user-defined type. At the program's entry point, initialize a complete instance of this composite type.
package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type employee struct {
	person
	tc       float64
	discount uint
}

func main() {
	myEmployee := employee{
		person: person{
			name: "Mike",
			age:  18,
		},
		tc:       2000,
		discount: 25,
	}
	fmt.Println(myEmployee)
}
