// Create a structure called Location and another called Person that contains Location as a field. Instantiate Person, filling in all the fields, and display the instance in the standard output. Then, change an internal field of Location and display the instance again to show the change.
package main

import "fmt"

type location struct {
	city    string
	country string
}

type person struct {
	name string
	age  uint8
	location
}

func main() {
	myAddress := location{"Tokyo", "Japan"}
	user := person{"mike", 18, myAddress}
	fmt.Println(user)

	user.country = "Brazil"
	user.city = "Sao Paulo"
	fmt.Println(user)
}
