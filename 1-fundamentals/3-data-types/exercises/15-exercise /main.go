// Create a function that receives a person's age and returns a boolean value indicating whether they are old enough to drink (age greater than or equal to 21). Then, store this value in a variable and display it in the main program.
package main

import (
	"fmt"
)

func isAdult(age byte) bool {
	return age >= 21
}
func main() {

	var canHeDrink bool = isAdult(20)
	fmt.Println(canHeDrink)

}
