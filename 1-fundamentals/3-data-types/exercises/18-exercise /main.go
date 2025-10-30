// Develop a function that accepts a text parameter and returns an integer type alias based on the length of the string.
package main

import (
	"fmt"
)

type myInt int

func StringLengthAlias(word string) myInt {
	return myInt(len(word))
}

func main() {
	result := StringLengthAlias("hello")
	fmt.Println(result)
}
