// Call a function that returns multiple values, but discard one of the results without storing it.
package main

import (
	"fmt"
)

func calc(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
func main() {
	sum, _ := calc(20, 30)
	fmt.Println(sum)
}
