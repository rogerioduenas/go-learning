// At the program's entry point, declare a base structure variable and assign values ​​to its fields.
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
