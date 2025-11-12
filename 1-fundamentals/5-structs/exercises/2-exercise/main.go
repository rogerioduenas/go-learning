// At the program's entry point, declare a variable with an explicit type of a created structure and assign individual values ​​to each field.
package main

import "fmt"

type student struct {
	course string
	campus string
}

func main() {
	var mike student
	mike.course = "Mathematics"
	mike.campus = "Tokyo"

	fmt.Println(mike)
	fmt.Println(mike.course)
	fmt.Println(mike.campus)
}
