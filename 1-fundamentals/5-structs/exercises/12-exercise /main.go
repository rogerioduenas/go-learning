// Create a function that receives a structure as an argument and displays one of its fields to standard output.
package main

import "fmt"

type person struct {
	name string
	age  uint8
}

func showPerson(stc person) {
	fmt.Println(stc.name)
}

func main() {
	user := person{"Mike", 21}
	showPerson(user)
}
