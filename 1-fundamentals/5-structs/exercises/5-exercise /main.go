// Create a structure called Location with two fields related to location information. Create a structure called Person and add Location as a field within it, representing a composition between types.
package main

import "fmt"

type location struct {
	country string
	city    string
}

type person struct {
	name string
	age  int8
	location
}

func main() {
	var user person
	user = person{"Mike", 18, location{"japan", "tokyo"}}
	fmt.Println(user)
}
