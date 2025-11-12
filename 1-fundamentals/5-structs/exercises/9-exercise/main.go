// Create a structure called Person and three instances of it, each initialized in a different way. Display each instance in the standard output, noting the differences between the initialization formats.
package main

import "fmt"

type person struct {
	name string
	age uint8
}

func main() {
	user1 := person{name: "Mike", age: 20}
	user2 := person{"John", 50}
	var user3 person
	user3.name = "Yuki"
	user3.age = 15

	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(user3)
}
