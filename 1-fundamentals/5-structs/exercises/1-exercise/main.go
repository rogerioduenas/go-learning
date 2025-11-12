// Create a structure composed of two simple fields: a text field and a small integer field.
package main

import "fmt"

type person struct {
	name string
	age  int8
}

func main() {
	fmt.Println(person{"Mike", 19})
}
