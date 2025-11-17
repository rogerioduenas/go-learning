// Create a function that takes a common numeric value and displays the result of a simple operation without modifying the original value.

package main

import "fmt"

func double(n1 int) int {
	return n1 + n1
}

func main() {
	var1 := 10

	fmt.Println(var1)
	fmt.Println(double(var1))
	fmt.Println(var1)
}
