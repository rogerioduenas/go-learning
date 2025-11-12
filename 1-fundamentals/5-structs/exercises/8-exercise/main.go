// Create a structure called Person. Create a new instance of this structure using named initialization to fill only a few specific fields.
package main

import "fmt"

type person struct {
	name string
	age uint8
	city string
	country string
}

func main() {
	user := person{name: "Mike", country: "Japan"}
	fmt.Println(user)
}
