// Create a structure with two fields. At the program's entry point, declare a variable of this structure, assign values ​​to both fields, and display the complete variable in the standard output.
package main

import "fmt"

type person struct {
	name string
	age  int8
}

func main() {
	var student person
	student = person{"Roger", 34}
	fmt.Println(student)
}
