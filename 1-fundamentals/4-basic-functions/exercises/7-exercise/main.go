// Create an anonymous function assigned to a variable, which receives a text argument and returns the same value after displaying it.
package main

import (
	"fmt"
)

var showText = func(text string) string {
	fmt.Println(text)
	return text
}

func main() {
	showText("Any text")
}
