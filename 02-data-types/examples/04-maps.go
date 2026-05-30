package main

import "fmt"

func main() {
	// Create a map
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Charlie"] = 92

	fmt.Println("Scores:", scores)
	fmt.Println("Alice's score:", scores["Alice"])

	// Map with initialization
	person := map[string]string{
		"name":  "John",
		"email": "john@example.com",
		"city":  "New York",
	}

	fmt.Println("Person:", person)

	// Check if key exists
	email, exists := person["email"]
	fmt.Println("Email exists:", exists, "Email:", email)

	unknown, exists := person["age"]
	fmt.Println("Age exists:", exists, "Age:", unknown)

	// Iterate over map
	fmt.Println("\nIterating over scores:")
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// Delete from map
	delete(scores, "Bob")
	fmt.Println("\nAfter deleting Bob:", scores)
}
