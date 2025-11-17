// Create a function that returns a newly allocated memory address containing an initial numeric value.

package main

import "fmt"

func newInt(initial int) *int {
	x := initial
	return &x
}

func main() {
	ptr := newInt(10)
	fmt.Println(ptr, *ptr)
}
