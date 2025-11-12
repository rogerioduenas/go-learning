// Create three nested structures: a main structure containing an inner structure, which in turn contains another inner structure. Instantiate all structures, assign values ​​to each field at all levels, and display the final result.
package main

import "fmt"

type city struct {
	name string
}

type address struct {
	street string
	city   city
}

type person struct {
	name    string
	address address
}

func main() {

	user := person{
		name: "Mike",
		address: address{
			street: "Any street",
			city: city{
				name: "Tokyo",
			},
		},
	}

	fmt.Println(user)
	fmt.Println(user.address.city.name)
	fmt.Println(user.address.street)
}
