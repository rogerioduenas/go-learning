// Create a variable and modify its value later without creating a new variable.
package main

import (
	"fmt"
)

func main() {
	var example string = "initial value"
	fmt.Println(example)
	example = "modified value"
	fmt.Println(example)
}