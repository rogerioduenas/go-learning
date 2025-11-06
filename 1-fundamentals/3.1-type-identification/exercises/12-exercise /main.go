// Create a function that receives a list containing values ​​of different types and displays the index and type of each item.
package main

import (
	"fmt"
	"reflect"
)

func checkTypes(list []any) {
	for index, value := range list {
		kind := reflect.TypeOf(value).Kind()
		fmt.Printf("Index %v -> %s\n", index, kind)
	}
}

func main() {
	list := []any{1, 2.0, true, [2]string{"a", "b"}}
	checkTypes(list)
}
