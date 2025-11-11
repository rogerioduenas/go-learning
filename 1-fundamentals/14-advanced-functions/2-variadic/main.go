package main

import "fmt"

func variadic(numbers ...int) {
	fmt.Println(numbers)
}

func sum(numbers ...int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func write(text string, numbers ...int) {
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	variadic(1, 2, 3)
	fmt.Println(sum(5, 5, 5))
	write("Hello world:", 1, 2, 3)
}
