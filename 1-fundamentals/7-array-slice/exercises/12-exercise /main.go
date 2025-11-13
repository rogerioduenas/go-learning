// Create a function that receives a dynamic collection of numbers and displays all of their values.

package main

import "fmt"

func showNumbersOfList(list []int) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func main() {
	slc := []int{1, 2, 3}
	showNumbersOfList(slc)
}
