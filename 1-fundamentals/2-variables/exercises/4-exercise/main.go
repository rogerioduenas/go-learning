// Create a constant representing a greeting and a variable representing a name, and combine both into a message displayed in the console.
package main

import (
	"fmt"
)

func main() {
	const greeting = "Good Morning"
	var name string = "Mike"
	fmt.Println(greeting, name)
}