// Implement a function that receives a heterogeneous list and uses a multi-case conditional structure to handle each data type differently.
package main

import (
	"errors"
	"fmt"
	"reflect"
)

func checkTypes(list any) error {
	val := reflect.ValueOf(list)

	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {

			item := val.Index(i).Interface()

			switch item.(type) {
			case int:
				fmt.Println("It's a int")
			case bool:
				fmt.Println("It's a bool")
			default:
				fmt.Println("Unsupported value")
			}
		}
		return nil
	}

	return errors.New("It isn't a list")
}

func main() {
	slc := []any{1, true, [1]int{}}
	checkTypes(slc)
}
