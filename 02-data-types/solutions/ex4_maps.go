package main

import "fmt"

func main() {
	// Create a map to store student grades
	grades := make(map[string]int)

	// Add students
	grades["Alice"] = 95
	grades["Bob"] = 87
	grades["Charlie"] = 92

	// Print all grades
	fmt.Println("Student Grades:")
	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}

	// Update a grade
	grades["Bob"] = 90
	fmt.Printf("\nAfter updating Bob's grade: %d\n", grades["Bob"])

	// Check if student exists
	score, exists := grades["Alice"]
	if exists {
		fmt.Printf("Alice's grade: %d (exists: %v)\n", score, exists)
	}

	unknownScore, exists := grades["Diana"]
	fmt.Printf("Diana's grade: %d (exists: %v)\n", unknownScore, exists)
}
