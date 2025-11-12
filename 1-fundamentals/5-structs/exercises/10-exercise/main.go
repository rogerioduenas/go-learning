// Create a structure called Address and another called Person that contains Address as a field. Create an Address variable and use it as the value of the corresponding field in Person.
package main

import "fmt"

type address struct {
	street string
	city string
}

type person struct {
	name string
	age uint8
	address
}

func main() {
	var myAddress = address{street: "any street", city: "Osaka"}
	user := person{name: "mike", age: 53, address: myAddress}
	fmt.Println(user)
}
