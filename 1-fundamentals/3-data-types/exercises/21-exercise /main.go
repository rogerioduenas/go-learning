// Create an error handler that receives an error, checks if it exists, and displays different messages depending on the condition.
package main

import (
	"errors"
	"fmt"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Error not found!")
}

func main() {
	checkError(errors.New("Something went wrong"))
	checkError(nil)
}
