package main

import "fmt"

// This function creates a closure that captures the variable `text`. The closure retains the value stored within the main function, even after the main function has finished.
func closure() func() {
	text := "Text inside CLOSURE function"

	function := func() {
		fmt.Println(text)
	}

	return function
}

// This function creates a closure that maintains the state of `number`. The variable continues to exist between calls, accumulating the previous value.
func closure2() func() int {
	number := 0

	function := func() int {
		number++
		return number
	}

	return function
}

func main() {
	text := "Text inside MAIN function"
	fmt.Println(text)

	newFunc := closure()
	newFunc()

	newFunc2 := closure2()
	fmt.Println(newFunc2())
	fmt.Println(newFunc2())
	fmt.Println(newFunc2())
}
