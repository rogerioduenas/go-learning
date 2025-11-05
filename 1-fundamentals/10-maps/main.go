package main

import "fmt"

func main() {

	fmt.Println("----------create----------")
	user := map[string]string{
		"firstName": "John",
		"lastName":  "Doe",
	}
	fmt.Println(user)
	fmt.Println(user["firstName"])
	fmt.Println("----------add item----------")
	user["city"] = "LA"
	fmt.Println(user)

	fmt.Println("----------nested map----------")
	user2 := map[string]map[string]string{
		"name": {
			"first": "Jane",
			"last":  "Doe",
		},
		"course": {
			"name":   "IT",
			"campus": "NY",
		},
	}
	fmt.Println(user2)
	fmt.Println(user2["course"]["name"])

	fmt.Println("----------delete----------")
	delete(user2, "course")
	fmt.Println(user2)

	fmt.Println("----------add item----------")
	user2["zodiac"] = map[string]string{
		"name": "taurus",
	}
	fmt.Println(user2)
	fmt.Println(user2["zodiac"]["name"])
}
