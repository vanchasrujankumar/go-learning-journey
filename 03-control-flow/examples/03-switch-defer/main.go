package main

import "fmt"

func dayType(day string) string {
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		return "Weekday"
	case "Saturday", "Sunday":
		return "Weekend"
	default:
		return "Invalid day"
	}
}

func main() {
	fmt.Println("=== Switch with Expression ===")
	days := []string{"Monday", "Saturday", "Funday"}
	for _, d := range days {
		fmt.Printf("%s: %s\n", d, dayType(d))
	}

	fmt.Println("\n=== Switch without Expression ===")
	score := 85
	var grade string
	switch {
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "F"
	}
	fmt.Printf("Score %d -> Grade %s\n", score, grade)

	fmt.Println("\n=== Defer Stacking ===")
	fmt.Println("Start")
	defer fmt.Println("First deferred")
	defer fmt.Println("Second deferred")
	defer fmt.Println("Third deferred")
	fmt.Println("End")
}
