// Create an anonymous function that generates and returns a custom error, and invoke it in the main program.
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := func() error {
		return errors.New("Any error")
	}()

	if err != nil {
		fmt.Println(err)
	}
}
