// Implement a function that receives a text and displays a message indicating that the value is a string.
package main

import (
	"errors"
	"fmt"
)

func isString(arg any) (string, error) {
	if _, ok := arg.(string); ok {
		return "It is a string", nil
	}
	return "", errors.New("It isn't a string.")
}

func main() {
	fmt.Println(isString("If you are reading this, it means you are very curious."))
	fmt.Println(isString([]int{1, 2, 3}))
}
