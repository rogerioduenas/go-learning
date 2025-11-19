package main

import "fmt"

// invertByValue returns a new number with its sign inverted.
// It does NOT modify the original value.
func invertByValue(n int) int {
	return n * -1
}

// invertByPointer inverts the sign of the number stored at the pointer.
// It DOES modify the original value in memory.
func invertByPointer(n *int) int {
	*n = *n * -1
	return *n
}

func main() {
	original := 20
	fmt.Println("Original number:", original)
	fmt.Println("Modified (copy):", invertByValue(original))

	fmt.Println("----------------")

	value := 40
	fmt.Println("Modified (pointer):", invertByPointer(&value))
	fmt.Println("Original value after pointer modification:", value)
}
