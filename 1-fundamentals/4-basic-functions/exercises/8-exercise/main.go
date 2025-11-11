// Create an anonymous function, store the return value in a variable, and display the returned value.
package main

import (
	"fmt"
)

func main() {
	result := func(text string) string {
		return text
	}("Any text")

	fmt.Println(result)
}
