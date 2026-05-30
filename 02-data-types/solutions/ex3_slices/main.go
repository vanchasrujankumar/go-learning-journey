package main

import "fmt"

func main() {
	// Create array and slice from it
	numbers := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := numbers[2:7] // Elements from index 2 to 6

	fmt.Println("Original slice:", slice)
	fmt.Println("Slice length:", len(slice))
	fmt.Println("Slice capacity:", cap(slice))

	// Append new elements
	slice = append(slice, 11, 12, 13)
	fmt.Println("\nAfter append:", slice)
	fmt.Println("New length:", len(slice))
	fmt.Println("New capacity:", cap(slice))

	// Modify elements
	slice[0] = 100
	fmt.Println("\nAfter modifying first element:", slice)

	// Create a new slice independently
	newSlice := []int{50, 60, 70}
	newSlice = append(newSlice, 80)
	fmt.Println("\nNew slice:", newSlice)
}
