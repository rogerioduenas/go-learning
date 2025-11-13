// Create a simple structure that represents a generic entity with two basic fields, one text and one numeric.
package main

import "fmt"

type person struct {
	name string
	age  uint8
}

func main() {

	user := person{
		name: "Mike",
		age:  18,
	}

	fmt.Println(user)
	fmt.Println(user.name)
	fmt.Println(user.age)

}
