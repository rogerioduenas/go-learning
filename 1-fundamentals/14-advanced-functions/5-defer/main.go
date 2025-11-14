package main

import "fmt"

func basicDefer() {
	defer fmt.Println("---with--- defer executes LAST")
	fmt.Println("---without--- defer executes FIRST")
}

// defer runs before the function exits
func isApproved(g1, g2 uint8) bool {
	defer fmt.Println("Did the student pass?")
	fmt.Println("Function execution starts here!!!")
	if (g1+g2)/2 > 5 {
		return true
	}
	return false
}

func executionOrder() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("third")
}

func main() {
	basicDefer()
	fmt.Println("---------------------")
	fmt.Println(isApproved(1, 3))
	fmt.Println("---------------------")
	executionOrder()
}
