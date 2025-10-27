package main

import (
	"fmt"
	"reflect"
)

// when you know the type you want to check
func isTypeInt(data interface{} /*|| any*/) {
	if value, ok := data.(int); ok {
		fmt.Println("It's a Int:", value)
	} else {
		fmt.Println("It isn't a Int")
	}
}

// when you don't know the type you're going to check
func whatTypeIsIt(data any /*|| interface{}*/) {
	myTypeIs := reflect.TypeOf(data).Kind()
	fmt.Println(myTypeIs)
}

// switch to check multiple types
func checkSomeTypes(data any) {
	switch value := data.(type) {
	case int:
		fmt.Println("It's a int:", value)
	case string:
		fmt.Println("It's a string:", value)
	default:
		fmt.Println("Another type")
	}
}

func main() {
	isTypeInt(2.0)
	whatTypeIsIt([3]int{1, 2, 3})
	checkSomeTypes(12.50)
}
