// Create a new function that, internally, uses the existing function to perform an additional calculation before returning the result.
package main

import (
	"fmt"
)

func sum() int {
	return 10 + 10
}

func div(n int) int {
	return sum() / n
}

func main() {
	result := div(2)
	fmt.Println(result)
}
