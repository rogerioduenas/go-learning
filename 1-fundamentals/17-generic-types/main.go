package main

import "fmt"

// Using empty interfaces (pre-generics style)
type anything interface{}
type genericKey interface{}

func printAnyOld(i interface{}) {
	fmt.Println(i)
}

func printLegacy(a anything) {
	fmt.Println(a)
}

// Generic type parameter T
func showAny[T any](value T) {
	fmt.Println("Received:", value)
}

// Generic function for slices
func showSlice[T any](slc []T) {
	fmt.Println("---- Showing slice ----")
	for i, v := range slc {
		fmt.Printf("[%d] => %v\n", i, v)
	}
	fmt.Println("------------------------")
}

// ---------------------------
//            main
// ---------------------------

func main() {

	// ---------- legacy generic-looking structures ----------
	m := map[genericKey]genericKey{
		1:            "apple",
		"banana":     true,
		true:         99.7,
		float64(80):  "water",
	}

	printAnyOld("hello world")
	printAnyOld(42)
	printAnyOld(false)

	fmt.Println("------------")

	printLegacy("legacy melon")
	printLegacy(777)
	printLegacy(true)

	fmt.Println("------------")
	fmt.Println(m)

	fmt.Println("\n==============================\n")

	// ---------- actual generics ----------
	showAny("hello worm")
	showAny(123)
	showAny(true)

	fmt.Println("-------------------------")

	// Generic slice examples
	fruits := []string{"onion", "garlic", "potato"}
	showSlice(fruits)

	numbers := []int{10, 20, 30, 40}
	showSlice(numbers)
}
