// Create a text variable, add a new word to it, and display the result.
package main

import (
	"fmt"
)

func main() {
	var hello string = "Hello"
	hello += " world"
	fmt.Println(hello)
}
