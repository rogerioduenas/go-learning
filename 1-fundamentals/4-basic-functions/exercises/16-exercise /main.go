// Create a control structure that chooses which function to execute based on a numerical condition, displaying the corresponding result.
package main

import (
	"fmt"
)

func canDrink() string {
	return "You are of legal age, you can drink."
}

func cannotDrink() string {
	return "You are underage, you cannot drink."
}

func checkAge(age int, can func() string, cannot func() string) string {

	if age >= 21 {
		return can()
	}

	return cannot()
}

func main() {
	fmt.Println(checkAge(21, canDrink, cannotDrink))
}
