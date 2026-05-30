package main

import "fmt"

func getGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	}
	return "F"
}

func main() {
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	if result := x * 2; result > 15 {
		fmt.Println("result is greater than 15")
	}

	scores := []int{95, 82, 73, 58, 100}
	for _, score := range scores {
		grade := getGrade(score)
		fmt.Printf("Score: %d -> Grade: %s\n", score, grade)
	}
}
