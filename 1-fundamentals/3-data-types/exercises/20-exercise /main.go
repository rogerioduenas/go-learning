// Implement a small system that stores multiple variables of different types and displays them all with formatted output that includes the name, value, and type of each variable.
package main

import (
	"fmt"
)

func identifyVars(args ...any) {
	for index, value := range args {
		fmt.Printf("My argument name is arg%d with the value %v and of type \"%T\" \n", index+1, value, value)
	}
}

func main() {
	identifyVars("a", []int{1, 2, 3}, true)
	fmt.Println()
}
