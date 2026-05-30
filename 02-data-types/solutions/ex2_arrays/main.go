package main

import "fmt"

func main() {
	// Create array of 5 numbers
	numbers := [5]int{10, 20, 30, 40, 50}

	// Print each element
	fmt.Println("Array elements:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index %d: %d\n", i, numbers[i])
	}

	// Calculate sum
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	fmt.Printf("\nSum of all elements: %d\n", sum)
	fmt.Printf("Array length: %d\n", len(numbers))
	fmt.Printf("Average: %.2f\n", float64(sum)/float64(len(numbers)))
}
