// Create a function that takes a boolean variable as input and returns a different message based on its value.
package main

import (
	"fmt"
)

func canIhaveDog(authorization bool) string{

	if authorization {
		return "You can have a dog."
	}
	return "You can't have a dog."
}

func main(){
	fmt.Println(canIhaveDog(true))
	fmt.Println(canIhaveDog(false))
}