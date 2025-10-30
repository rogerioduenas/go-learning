// Create multiple variables of different types within the same function and display their formatted values ​​in a single output.
package main

import (
	"fmt"
)

func studentData(name string, age byte, grade float32) {
	fmt.Printf("The student %s is %d years old and scored %.2f\n", name, age, grade)
}
func main() {
	studentData("Mike", 12, 8.7)
}
