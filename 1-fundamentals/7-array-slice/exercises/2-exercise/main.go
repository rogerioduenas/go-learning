// Create another collection of the same type, but storing text, and assign a value only to the first position.
package main

import "fmt"

func main() {
	array := [3]string{}
	array[0] = "first position"
	fmt.Println(array)
}
