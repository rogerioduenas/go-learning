// Create an error variable containing a custom message and display it.
package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error = errors.New("Any error")
	fmt.Println(err)
}
