package main

import "fmt"

func fibonacci(position uint) (result uint) {
	if position <= 1 {
		return position
	}
	result = fibonacci(position-2) + fibonacci(position-1)
	return
}

func main() {
	fmt.Println(fibonacci(7))
	fmt.Println("-----------------")

	position := uint(7)
	for i := uint(0); i <= position; i++{
		fmt.Println(fibonacci(i))
	}
}
