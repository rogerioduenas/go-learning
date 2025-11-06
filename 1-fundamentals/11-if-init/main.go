package main

import "fmt"

func main() {
	age := 21

	if ageToCheck := age; ageToCheck >= 18 {
		fmt.Println("Adult")
	} else if ageToCheck >= 12 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Child")
	}
}
