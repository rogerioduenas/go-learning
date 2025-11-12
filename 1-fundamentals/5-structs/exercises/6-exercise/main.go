// Create a structure called Location and another called Person that contains Location as a field. Assign values ​​to the internal Location field within Person and display the complete result.
package main

import "fmt"

type location struct {
	country string
	city    string
}

type person struct {
	name string
	age  uint8
	location
}

func main() {
	var student person
	student.country = "Japan"
	student.city = "Tokyo"
	//or
	// student = person{location: location{"Japan", "Tokyo"}}
	fmt.Println(student)
}