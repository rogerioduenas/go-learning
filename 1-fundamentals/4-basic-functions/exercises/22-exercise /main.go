// Create an anonymous function stored in a variable that returns another anonymous function. Call the inner function and display the result.
package main

import "fmt"

var fn = func() func(string) string {
	return func(text string) string {
		return text
	}
}

func main() {
	fmt.Println(fn()("any text"))
}
