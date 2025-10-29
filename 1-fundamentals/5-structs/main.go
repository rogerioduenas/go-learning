package main

import "fmt"

type user struct {
	name    string
	age     uint8
	gender  string
	address address
}

type address struct {
	street string
	number uint16
}

func main() {

	var user1 user

	user1.name = "Mike"
	user1.age = 18
	fmt.Println(user1)

	anyAddress := address{"any street", 0}
	user1.address = anyAddress
	fmt.Println(user1)

	user2 := user{"Kate", 21, "female", anyAddress}
	// user2 := user{"Kate", 21, "female"} error........
	// Therefore, it is necessary to assign all the fields of the struct.
	fmt.Println(user2)

	user3 := user{name: "Nick", age: 12}
	//To assign values ​​to only some fields of the struct, you can do it like this.
	fmt.Println(user3)
}
