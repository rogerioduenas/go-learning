// Develop a function that receives another function as a parameter and invokes it internally, displaying the returned result.
package main

import (
	"fmt"
)

func showText(text string) string {
	return text
}

func callback(function func(string) string, text string) string {
	return function(text)
}

func main() {
	fmt.Println(callback(showText, "Any text"))
}
