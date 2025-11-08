package main

import "fmt"

func main() {

	// simple for loop
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementing I", i)
	}

	fmt.Println("----------------------")

	// classic for loop
	for j := 0; j < 10; j++ {
		fmt.Println("Incrementing J", j)
	}

	fmt.Println("----------------------")

	// iterate over an array / slice
	names := [3]string{"John", "Jane", "Mike"}
	for index, name := range names {
		fmt.Println(index, name)
	}

	fmt.Println("----------------------")

	// show only the index
	names2 := [3]string{"John", "Jane", "Mike"}
	for index := range names2 {
		fmt.Println(index)
	}

	fmt.Println("----------------------")

	// ignore index using underscore and print only names
	names3 := [3]string{"John", "Jane", "Mike"}
	for _, name := range names3 {
		fmt.Println(name)
	}

	fmt.Println("----------------------")

	// iterate through a word and show the rune code
	for index, letter := range "WORD" {
		fmt.Println(index, letter)
	}

	fmt.Println("----------------------")

	// iterate through a word and show each letter
	for index, letter := range "WORD" {
		fmt.Println(index, string(letter))
	}

	fmt.Println("----------------------")

	// iterate through a map
	user := map[string]string{
		"firstName": "John",
		"lastName":  "Doe",
	}
	for key, value := range user {
		fmt.Println(key, value)
	}
}
