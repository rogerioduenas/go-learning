package main

import "fmt"

// With parameter
func gradeLevel(score int) string {
	switch score {
	case 10:
		return "Perfect score"
	case 9:
		return "Excellent"
	default:
		return "Needs improvement"
	}
}

// Without parameter
func checkTemperature(temp int) string {
	switch {
	case temp < 0:
		return "Below freezing"
	case temp >= 0:
		return "Above freezing"
	}
	return ""
}

// Calling a function inside a case
func getCategory() string {
	level := 2
	switch {
	case compute(level) == 2:
		return "Intermediate"
	case compute(level) == 3:
		return "Advanced"
	default:
		return "Unknown"
	}
}

func compute(x int) int {
	return x + 1
}

// With multiple items in a single case
func dayType(day string) string {
	switch day {
	case "Saturday", "Sunday":
		return "Weekend"
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "Weekday"
	default:
		return "Invalid day"
	}
}

// Using fallthrough
func checkSpeed(speed int) string {
	switch {
	case speed < 60:
		return "Normal speed"
	case speed >= 60:
		fmt.Println("Warning: fast!")
		fallthrough // Continues to next case manually
	case speed > 100:
		return "Too fast!"
	}
	return ""
}

func main() {
	fmt.Println(gradeLevel(9))
	fmt.Println(checkTemperature(-5))
	fmt.Println(getCategory())
	fmt.Println(dayType("Saturday"))
	fmt.Println(checkSpeed(80))
}
