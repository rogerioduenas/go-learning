// Create a structure called `location` with two fields related to location information (for example, `city` and `country`). At the program's entry point (main), declare a variable of this structure, assign values ​​to both fields, and display the complete variable in the standard output.
package main

import "fmt"

type location struct {
	country string
	city string
}

func main() {
	var address location
	address = location{country: "Japan", city: "Tokyo"}
	fmt.Println(address)
}
