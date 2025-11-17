// Develop a code that stores multiple memory addresses in a collection and modifies each value pointed to.

package main

import "fmt"

func modifyList(slc []*int) {
	for _, value := range slc {
		*value = 0
	}
}

func main() {
	var1, var2, var3 := 10, 20, 30
	list := []*int{&var1, &var2, &var3}

	fmt.Println(var1, var2, var3)

	modifyList(list)

	fmt.Println(var1, var2, var3)
}
